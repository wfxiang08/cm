package proxy

import (
	"sync"

	"github.com/juju/errors"
	"github.com/wandoulabs/cm/config"
	"github.com/wandoulabs/cm/mysql"
)

const (
	Master = "master"
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
