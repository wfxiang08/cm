package proxy

import (
	"github.com/juju/errors"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/sqlparser"
	"strings"
)

var nstring = sqlparser.String

func (c *Conn) handleSet(stmt *sqlparser.Set, sql string) error {
	switch stmt.Scope {
	case "global":
		log.Warningf("set global, %s", sql)
		return errors.Errorf("set global not allowed, %s", sql)
	case "session":
		log.Warning("set session")
	}

	k := string(stmt.Exprs[0].Name.Name)

	switch strings.ToUpper(k) {
	case `AUTOCOMMIT`:
		return c.handleSetAutoCommit(stmt.Exprs[0].Expr, sql)
	case `NAMES`:
		return c.handleSetNames(stmt.Exprs[0].Expr)
	default:
		//todo:strict condition
		return c.handleShow(nil, sql, nil) //errors.Errorf("set %s is not supported now", k)
	}
}

func (c *Conn) handleSetAutoCommit(val sqlparser.ValExpr, sql string) error {
	value, ok := val.(sqlparser.NumVal)
	if !ok {
		return errors.Errorf("set autocommit error")
	}

	switch value[0] {
	case '1':
		//default value is 1, no need to do anything
		log.Warning("set autocommit 1")
	case '0':
		log.Warning("set autocommit 0")
		c.server.IncCounter("set autocommit 0")
		c.status &= ^mysql.SERVER_STATUS_AUTOCOMMIT
	default:
		return errors.Errorf("invalid autocommit flag %s", value)
	}

	err := c.writeOkFlush(nil)
	return errors.Trace(err)
}

func (c *Conn) handleSetNames(val sqlparser.ValExpr) error {
	value, ok := val.(sqlparser.StrVal)
	if !ok {
		return errors.Errorf("set names charset error")
	}

	charset := strings.ToLower(string(value))
	cid, ok := mysql.CharsetIds[charset]
	if !ok {
		return errors.Errorf("invalid charset %s", charset)
	}

	c.charset = charset
	c.collation = cid

	err := c.writeOkFlush(nil)
	return errors.Trace(err)
}
