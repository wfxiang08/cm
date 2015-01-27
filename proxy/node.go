package proxy

import (
	"github.com/juju/errors"
	"github.com/wandoulabs/cm/config"
	"github.com/wandoulabs/cm/mysql"
)

const (
	Master = "master"
)

type Shard struct {
	cfg    config.ShardConfig
	master *mysql.DB
}

func (shard *Shard) String() string {
	return shard.cfg.Id
}

func (shard *Shard) Close() {
	shard.master.Close()
}

func (shard *Shard) getMasterConn() (*mysql.SqlConn, error) {
	db := shard.master
	if db == nil {
		return nil, errors.Errorf("master is down")
	}

	return db.GetConn()
}

func (shard *Shard) openDB(addr string) (*mysql.DB, error) {
	db, err := mysql.Open(addr, shard.cfg.User, shard.cfg.Password, "")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConnNum(100)

	return db, nil
}
