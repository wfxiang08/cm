package proxy

import (
	"io"
	"net"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/juju/errors"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/config"
	"github.com/wandoulabs/cm/vt/schema"
	"github.com/wandoulabs/cm/vt/tabletserver"
)

type Server struct {
	configFile  string
	cfg         *config.Config
	addr        string
	user        string
	password    string
	running     int32
	listener    net.Listener
	nodes       map[string]*Node
	schemas     map[string]*Schema
	autoSchamas map[string]*tabletserver.SchemaInfo
	rwlock      sync.RWMutex
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

func (s *Server) loadSchemaInfo() error {
	if err := s.parseNodes(); err != nil {
		return errors.Trace(err)
	}

	if err := s.parseSchemas(); err != nil {
		return errors.Trace(err)
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

	return nil
}

func makeServer(configFile string) *Server {
	cfg, err := config.ParseConfigFile(configFile)
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	s := &Server{
		configFile:  configFile,
		cfg:         cfg,
		addr:        cfg.Addr,
		user:        cfg.User,
		password:    cfg.Password,
		autoSchamas: make(map[string]*tabletserver.SchemaInfo),
	}

	return s
}

func NewServer(configFile string) (*Server, error) {
	s := makeServer(configFile)
	s.loadSchemaInfo()

	netProto := "tcp"
	if strings.Contains(netProto, "/") {
		netProto = "unix"
	}

	var err error
	s.listener, err = net.Listen(netProto, s.addr)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Infof("Server run MySql Protocol Listen(%s) at [%s]", netProto, s.addr)
	return s, nil
}

func (s *Server) cleanup() {
	for _, si := range s.autoSchamas {
		si.Close()
	}
}

func (s *Server) resetSchemaInfo() {
	s.cleanup()
	s.autoSchamas = make(map[string]*tabletserver.SchemaInfo)
	s.nodes = nil
	s.schemas = nil

	cfg, err := config.ParseConfigFile(s.configFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	s.cfg = cfg
	s.loadSchemaInfo()
}

func (s *Server) HandleReload(w http.ResponseWriter, req *http.Request) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	s.resetSchemaInfo()

	io.WriteString(w, "ok")
}

func (s *Server) Run() error {
	atomic.StoreInt32(&s.running, 1)

	for atomic.LoadInt32(&s.running) == 1 {
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
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	atomic.StoreInt32(&s.running, 0)
	if s.listener != nil {
		s.listener.Close()
	}

	s.cleanup()
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
