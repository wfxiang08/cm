package proxy

import (
	"sync"

	"github.com/juju/errors"
	"github.com/wandoulabs/cm/config"
	"github.com/wandoulabs/cm/mysql"
)

const (
	Master = "master"
	Slave  = "slave"
)

type Node struct {
	sync.Mutex
	server *Server
	cfg    config.NodeConfig
	master *mysql.DB
}

func (n *Node) String() string {
	return n.cfg.Name
}

func (n *Node) Close() {
	n.master.Close()
}

func (n *Node) getMasterConn() (*mysql.SqlConn, error) {
	n.Lock()
	db := n.master
	n.Unlock()

	if db == nil {
		return nil, errors.Errorf("master is down")
	}

	return db.GetConn()
}

func (n *Node) openDB(addr string) (*mysql.DB, error) {
	db, err := mysql.Open(addr, n.cfg.User, n.cfg.Password, "")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConnNum(n.cfg.IdleConns)
	return db, nil
}

func (s *Server) getNode(name string) *Node {
	return s.nodes[name]
}

func (s *Server) getNodeNames() []string {
	ret := make([]string, 0, len(s.nodes))
	for name, _ := range s.nodes {
		ret = append(ret, name)
	}

	return ret
}

func (s *Server) parseNodes() error {
	cfg := s.cfg
	s.nodes = make(map[string]*Node, len(cfg.Nodes))

	for _, v := range cfg.Nodes {
		if _, ok := s.nodes[v.Name]; ok {
			return errors.Errorf("duplicate node [%s].", v.Name)
		}

		n, err := s.parseNode(v)
		if err != nil {
			return errors.Trace(err)
		}

		s.nodes[v.Name] = n
	}

	return nil
}

func (s *Server) parseNode(cfg config.NodeConfig) (*Node, error) {
	n := &Node{
		server: s,
		cfg:    cfg,
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
