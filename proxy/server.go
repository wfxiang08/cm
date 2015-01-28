package proxy

import (
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/juju/errors"
	"github.com/ngaut/arena"
	stats "github.com/ngaut/gostats"
	log "github.com/ngaut/logging"
	"github.com/ngaut/tokenlimiter"
	"github.com/wandoulabs/cm/config"
	"github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/vt/tabletserver"
)

var (
	baseConnId uint32 = 10000
)

type Server struct {
	configFile        string
	cfg               *config.Config
	addr              string
	user              string
	password          string
	listener          net.Listener
	shards            map[string]*Shard
	schemas           map[string]*Schema
	autoSchamas       map[string]*tabletserver.SchemaInfo
	rwlock            *sync.RWMutex
	taskQ             chan *execTask
	concurrentLimiter *tokenlimiter.TokenLimiter

	counter *stats.Counters

	clients map[uint32]*Conn
}

type IServer interface {
	GetSchema(string) *Schema
	GetRowCacheSchema(string) (*tabletserver.SchemaInfo, bool)
	CfgGetPwd() string
	SkipAuth() bool
	GetToken() *tokenlimiter.Token
	ReleaseToken(token *tokenlimiter.Token)
	GetRWlock() *sync.RWMutex
	GetShard(shardId string) *Shard
	GetShardIds() []string
	AsynExec(task *execTask)
	IncCounter(key string)
	DecCounter(key string)
}

func (s *Server) IncCounter(key string) {
	s.counter.Add(key, 1)
}

func (s *Server) DecCounter(key string) {
	s.counter.Add(key, -1)
}

func (s *Server) GetToken() *tokenlimiter.Token {
	return s.concurrentLimiter.Get()
}

func (s *Server) ReleaseToken(token *tokenlimiter.Token) {
	s.concurrentLimiter.Put(token)
}

func (s *Server) GetShard(shardId string) *Shard {
	return s.shards[shardId]
}

func (s *Server) GetRowCacheSchema(db string) (*tabletserver.SchemaInfo, bool) {
	si, ok := s.autoSchamas[db]
	return si, ok
}

func (s *Server) GetShardIds() []string {
	ids := make([]string, 0, len(s.shards))
	for id, _ := range s.shards {
		ids = append(ids, id)
	}

	return ids
}

func (s *Server) parseShards() error {
	cfg := s.cfg
	s.shards = make(map[string]*Shard, len(cfg.Shards))

	for _, v := range cfg.Shards {
		if _, ok := s.shards[v.Id]; ok {
			return errors.Errorf("duplicate node [%s].", v.Id)
		}

		n, err := s.parseShard(v)
		if err != nil {
			return errors.Trace(err)
		}

		s.shards[v.Id] = n
	}

	return nil
}

func (s *Server) parseShard(cfg config.ShardConfig) (*Shard, error) {
	n := &Shard{
		cfg: cfg,
	}
	if len(cfg.Master) == 0 {
		return nil, errors.Errorf("must setting master MySQL node.")
	}

	var err error
	if n.master, err = n.openDB(cfg.Master); err != nil {
		return nil, errors.Trace(err)
	}

	return n, nil
}

func (s *Server) newConn(co net.Conn) *Conn {
	log.Info("newConn", co.RemoteAddr().String())
	c := &Conn{
		c:            co,
		pkg:          mysql.NewPacketIO(co),
		server:       s,
		connectionId: atomic.AddUint32(&baseConnId, 1),
		status:       mysql.SERVER_STATUS_AUTOCOMMIT,
		collation:    mysql.DEFAULT_COLLATION_ID,
		charset:      mysql.DEFAULT_CHARSET,
		alloc:        arena.NewArenaAllocator(32 * 1024),
		txConns:      make(map[string]*mysql.SqlConn),
	}
	c.salt, _ = mysql.RandomBuf(20)

	return c
}

func (s *Server) GetRWlock() *sync.RWMutex {
	return s.rwlock
}

func (s *Server) AsynExec(task *execTask) {
	s.taskQ <- task
}

func (s *Server) SkipAuth() bool {
	return s.cfg.SkipAuth
}

func (s *Server) CfgGetPwd() string {
	return s.cfg.Password
}

func (s *Server) loadSchemaInfo() error {
	if err := s.parseShards(); err != nil {
		return errors.Trace(err)
	}

	if err := s.parseSchemas(); err != nil {
		return errors.Trace(err)
	}

	for _, v := range s.cfg.Schemas {
		rc := v.RouterConifg
		var overrides []tabletserver.SchemaOverride
		for _, sc := range rc.TableRule {
			or := tabletserver.SchemaOverride{Name: sc.Table}
			pks := strings.Split(sc.ShardingKey, ",")
			for _, pk := range pks {
				or.PKColumns = append(or.PKColumns, strings.TrimSpace(pk))
			}
			or.Cache = &tabletserver.OverrideCacheDesc{Type: sc.RowCacheType, Prefix: or.Name, Table: or.Name}
			overrides = append(overrides, or)
		}

		//fix hard code node
		sc := s.cfg.Shards[0]
		s.autoSchamas[v.DB] = tabletserver.NewSchemaInfo(s.cfg.RowCacheConf, s.cfg.Shards[0].Master, sc.User, sc.Password, v.DB, overrides)
	}

	return nil
}

func makeServer(configFile string) *Server {
	cfg, err := config.ParseConfigFile(configFile)
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	log.Warningf("%#v", cfg)

	s := &Server{
		configFile:        configFile,
		cfg:               cfg,
		addr:              cfg.Addr,
		user:              cfg.User,
		password:          cfg.Password,
		autoSchamas:       make(map[string]*tabletserver.SchemaInfo),
		taskQ:             make(chan *execTask, 100),
		concurrentLimiter: tokenlimiter.NewTokenLimiter(100),
		counter:           stats.NewCounters("stats"),
		rwlock:            &sync.RWMutex{},
		clients:           make(map[uint32]*Conn),
	}

	f := func(wg *sync.WaitGroup, rs []interface{}, i int, co *mysql.SqlConn, sql string, args []interface{}) {
		r, err := co.Execute(sql, args...)
		if err != nil {
			log.Warning(err)
			rs[i] = err
		} else {
			rs[i] = r
		}

		wg.Done()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for task := range s.taskQ {
				f(task.wg, task.rs, task.idx, task.co, task.sql, task.args)
			}
		}()
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

func (s *Server) resetSchemaInfo() error {
	for _, c := range s.clients {
		if len(c.txConns) > 0 {
			return errors.Errorf("transaction exist")
		}
	}

	s.cleanup()
	s.autoSchamas = make(map[string]*tabletserver.SchemaInfo)
	for _, n := range s.shards {
		n.Close()
	}

	s.shards = nil
	s.schemas = nil

	cfg, err := config.ParseConfigFile(s.configFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Warningf("%#v", cfg)

	log.SetLevelByString(cfg.LogLevel)

	s.cfg = cfg
	s.loadSchemaInfo()
	return nil
}

func (s *Server) HandleReload(w http.ResponseWriter, req *http.Request) {
	log.Warning("trying to reload config")
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	if err := s.resetSchemaInfo(); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	io.WriteString(w, "ok")
}

func (s *Server) Run() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Errorf("accept error %s", err.Error())
			return err
		}

		go s.onConn(conn)
	}

	return nil
}

func (s *Server) Close() {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	if s.listener != nil {
		s.listener.Close()
		s.listener = nil
	}

	s.cleanup()
}

func (s *Server) onConn(c net.Conn) {
	conn := s.newConn(c)
	if err := conn.Handshake(); err != nil {
		log.Errorf("handshake error %s", errors.ErrorStack(err))
		c.Close()
		return
	}

	const key = "connections"

	s.IncCounter(key)
	defer func() {
		s.DecCounter(key)
		log.Infof("close %s", conn)
	}()

	s.rwlock.Lock()
	s.clients[conn.connectionId] = conn
	s.rwlock.Unlock()

	conn.Run()
}
