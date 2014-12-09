package proxy

import (
	"net"
	"runtime"
	"strings"

	"github.com/juju/errors"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/config"
	"github.com/wandoulabs/cm/vt/schema"
	"github.com/wandoulabs/cm/vt/tabletserver"
)

type Server struct {
	cfg         *config.Config
	addr        string
	user        string
	password    string
	running     bool
	listener    net.Listener
	nodes       map[string]*Node
	schemas     map[string]*Schema
	autoSchamas map[string]*tabletserver.SchemaInfo
}

func GetRowCacheType(rowCacheType string) int {
	switch rowCacheType {
	case "RW":
		return schema.CACHE_RW
	case "W":
		return schema.CACHE_W
	default:
		return schema.CACHE_NONE
	}
}

func NewServer(cfg *config.Config) (*Server, error) {
	s := &Server{
		cfg:         cfg,
		addr:        cfg.Addr,
		user:        cfg.User,
		password:    cfg.Password,
		autoSchamas: make(map[string]*tabletserver.SchemaInfo),
	}

	if err := s.parseNodes(); err != nil {
		return nil, errors.Trace(err)
	}

	if err := s.parseSchemas(); err != nil {
		return nil, errors.Trace(err)
	}

	for _, v := range s.cfg.Schemas {
		rc := v.RulesConifg
		var overrides []tabletserver.SchemaOverride
		for _, sc := range rc.ShardRule {
			or := tabletserver.SchemaOverride{Name: sc.Table}
			pks := strings.Split(sc.Key, ",")
			for _, pk := range pks {
				or.PKColumns = append(or.PKColumns, strings.TrimSpace(pk))
			}
			or.Cache = &tabletserver.OverrideCacheDesc{Type: sc.RowCacheType, Prefix: or.Name, Table: or.Name}
			overrides = append(overrides, or)
		}

		//fix hard code node
		s.autoSchamas[v.DB] = tabletserver.NewSchemaInfo(128*1024*1024, s.cfg.Nodes[0].Master, s.cfg.User, s.cfg.Password, v.DB, overrides)
	}

	var err error
	netProto := "tcp"
	if strings.Contains(netProto, "/") {
		netProto = "unix"
	}
	s.listener, err = net.Listen(netProto, s.addr)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Infof("Server run MySql Protocol Listen(%s) at [%s]", netProto, s.addr)
	return s, nil
}

func (s *Server) Run() error {
	s.running = true

	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Errorf("accept error %s", err.Error())
			continue
		}

		go s.onConn(conn)
	}

	return nil
}

func (s *Server) Close() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *Server) onConn(c net.Conn) {
	conn := s.newConn(c)

	defer func() {
		if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Errorf("onConn panic %v: %v\n%s", c.RemoteAddr().String(), err, buf)
		}

		conn.Close()
	}()

	if err := conn.Handshake(); err != nil {
		log.Errorf("handshake error %s", err.Error())
		c.Close()
		return
	}

	conn.Run()

}
