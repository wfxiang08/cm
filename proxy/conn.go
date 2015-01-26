package proxy

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"runtime"
	"strings"

	"github.com/juju/errors"
	"github.com/ngaut/arena"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/hack"
	"github.com/wandoulabs/cm/mysql"
)

var DEFAULT_CAPABILITY uint32 = mysql.CLIENT_LONG_PASSWORD | mysql.CLIENT_LONG_FLAG |
	mysql.CLIENT_CONNECT_WITH_DB | mysql.CLIENT_PROTOCOL_41 |
	mysql.CLIENT_TRANSACTIONS | mysql.CLIENT_SECURE_CONNECTION

//client <-> proxy
type Conn struct {
	pkg          *mysql.PacketIO
	c            net.Conn
	server       IServer
	capability   uint32
	connectionId uint32
	status       uint16
	collation    mysql.CollationId
	charset      string
	user         string
	db           string
	salt         []byte
	lastInsertId int64
	affectedRows int64
	alloc        arena.ArenaAllocator
	txConns      map[string]*mysql.SqlConn
	lastCmd      string
}

func (c *Conn) String() string {
	return fmt.Sprintf("conn: %s, status: %d, charset: %s, user: %s, db: %s, lastInsertId: %d",
		c.c.RemoteAddr(), c.status, c.charset, c.user, c.db, c.lastInsertId,
	)
}

func (c *Conn) schema() *Schema {
	return c.server.GetSchema(c.db)
}

func (c *Conn) Handshake() error {
	if err := c.writeInitialHandshake(); err != nil {
		return errors.Trace(err)
	}

	c.flush()

	if err := c.readHandshakeResponse(); err != nil {
		c.writeError(err)
		return errors.Trace(err)
	}

	err := c.writeOkFlush(nil)
	c.pkg.Sequence = 0

	return err
}

func (c *Conn) Close() error {
	c.rollback()
	c.c.Close()

	return nil
}

func (c *Conn) writeInitialHandshake() error {
	data := make([]byte, 4, 128)

	//min version 10
	data = append(data, 10)
	//server version[00]
	data = append(data, mysql.ServerVersion...)
	data = append(data, 0)
	//connection id
	data = append(data, byte(c.connectionId), byte(c.connectionId>>8), byte(c.connectionId>>16), byte(c.connectionId>>24))
	//auth-plugin-data-part-1
	data = append(data, c.salt[0:8]...)
	//filter [00]
	data = append(data, 0)
	//capability flag lower 2 bytes, using default capability here
	data = append(data, byte(DEFAULT_CAPABILITY), byte(DEFAULT_CAPABILITY>>8))
	//charset, utf-8 default
	data = append(data, uint8(mysql.DEFAULT_COLLATION_ID))
	//status
	data = append(data, byte(c.status), byte(c.status>>8))
	//below 13 byte may not be used
	//capability flag upper 2 bytes, using default capability here
	data = append(data, byte(DEFAULT_CAPABILITY>>16), byte(DEFAULT_CAPABILITY>>24))
	//filter [0x15], for wireshark dump, value is 0x15
	data = append(data, 0x15)
	//reserved 10 [00]
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	//auth-plugin-data-part-2
	data = append(data, c.salt[8:]...)
	//filter [00]
	data = append(data, 0)

	return c.writePacket(data)
}

func (c *Conn) readPacket() ([]byte, error) {
	return c.pkg.ReadPacket()
}

func (c *Conn) writePacket(data []byte) error {
	return c.pkg.WritePacket(data)
}

func (c *Conn) flush() error {
	return c.pkg.Flush()
}

func (c *Conn) readHandshakeResponse() error {
	data, err := c.readPacket()

	if err != nil {
		return errors.Trace(err)
	}

	pos := 0
	//capability
	c.capability = binary.LittleEndian.Uint32(data[:4])
	pos += 4
	//skip max packet size
	pos += 4
	//charset, skip, if you want to use another charset, use set names
	c.collation = mysql.CollationId(data[pos])
	pos++
	//skip reserved 23[00]
	pos += 23
	//user name
	c.user = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
	pos += len(c.user) + 1
	//auth length and auth
	authLen := int(data[pos])
	pos++
	auth := data[pos : pos+authLen]
	checkAuth := mysql.CalcPassword(c.salt, []byte(c.server.CfgGetPwd()))
	if !bytes.Equal(auth, checkAuth) && !c.server.CfgIsSkipAuth() {
		return errors.Trace(mysql.NewDefaultError(mysql.ER_ACCESS_DENIED_ERROR, c.c.RemoteAddr().String(), c.user, "Yes"))
	}

	pos += authLen
	if c.capability|mysql.CLIENT_CONNECT_WITH_DB > 0 {
		if len(data[pos:]) == 0 {
			return nil
		}

		db := string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(c.db) + 1

		if err := c.useDB(db); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

func (c *Conn) Run() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]

			log.Errorf("lastCmd %s, %v, %s", c.lastCmd, err, buf)
		}

		c.Close()
	}()

	for {
		c.alloc.Reset()
		data, err := c.readPacket()
		if err != nil {
			if err.Error() != io.EOF.Error() {
				log.Info(err)
			}
			return
		}

		if err := c.dispatch(data); err != nil {
			log.Errorf("dispatch error %s, %s", errors.ErrorStack(err), c)
			if err != mysql.ErrBadConn { //todo: fix this
				c.writeError(err)
			}
		}

		c.pkg.Sequence = 0
	}
}

func (c *Conn) dispatch(data []byte) error {
	cmd := data[0]
	data = data[1:]

	log.Debug(c.connectionId, cmd, hack.String(data))
	c.lastCmd = hack.String(data)

	token := c.server.GetToken()

	c.server.GetRWlock().RLock()
	defer func() {
		c.server.GetRWlock().RUnlock()
		c.server.ReleaseToken(token)
	}()

	c.server.IncCounter(mysql.MYSQL_COMMAND(cmd).String())

	switch mysql.MYSQL_COMMAND(cmd) {
	case mysql.COM_QUIT:
		c.Close()
		return nil
	case mysql.COM_QUERY:
		return c.handleQuery(hack.String(data))
	case mysql.COM_PING:
		return c.writeOkFlush(nil)
	case mysql.COM_INIT_DB:
		log.Debug(cmd, hack.String(data))
		if err := c.useDB(hack.String(data)); err != nil {
			return errors.Trace(err)
		}

		return c.writeOkFlush(nil)
	case mysql.COM_FIELD_LIST:
		return c.handleFieldList(data)
	case mysql.COM_STMT_PREPARE:
		// not support server side prepare yet
	case mysql.COM_STMT_EXECUTE:
		log.Fatal("not support", data)
	case mysql.COM_STMT_CLOSE:
		return c.handleStmtClose(data)
	case mysql.COM_STMT_SEND_LONG_DATA:
		log.Fatal("not support", data)
	case mysql.COM_STMT_RESET:
		log.Fatal("not support", data)
	default:
		msg := fmt.Sprintf("command %d not supported now", cmd)
		return mysql.NewError(mysql.ER_UNKNOWN_ERROR, msg)
	}

	return nil
}

func (c *Conn) useDB(db string) error {
	db = strings.ToLower(db)
	if s := c.server.GetSchema(db); s == nil {
		return mysql.NewDefaultError(mysql.ER_BAD_DB_ERROR, db)
	} else {
		c.db = db
	}

	return nil
}

func (c *Conn) writeOkFlush(r *mysql.Result) error {
	if err := c.writeOK(r); err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(c.flush())
}

func (c *Conn) writeOK(r *mysql.Result) error {
	if r == nil {
		r = &mysql.Result{Status: c.status}
	}
	data := c.alloc.AllocBytesWithLen(4, 32)
	data = append(data, mysql.OK_HEADER)
	data = append(data, mysql.PutLengthEncodedInt(r.AffectedRows)...)
	data = append(data, mysql.PutLengthEncodedInt(r.InsertId)...)
	if c.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(r.Status), byte(r.Status>>8))
		data = append(data, 0, 0)
	}

	err := c.writePacket(data)
	if err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(c.flush())
}

func (c *Conn) writeError(e error) error {
	var m *mysql.SqlError
	var ok bool
	if m, ok = e.(*mysql.SqlError); !ok {
		m = mysql.NewError(mysql.ER_UNKNOWN_ERROR, e.Error())
	}

	data := make([]byte, 4, 16+len(m.Message))
	data = append(data, mysql.ERR_HEADER)
	data = append(data, byte(m.Code), byte(m.Code>>8))
	if c.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, '#')
		data = append(data, m.State...)
	}

	data = append(data, m.Message...)

	err := c.writePacket(data)
	if err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(c.flush())
}

func (c *Conn) writeEOF(status uint16) error {
	data := c.alloc.AllocBytesWithLen(4, 9)

	data = append(data, mysql.EOF_HEADER)
	if c.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}

	err := c.writePacket(data)
	return errors.Trace(err)
}
