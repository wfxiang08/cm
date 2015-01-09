package mysql

import (
	"container/list"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/juju/errors"
)

type DB struct {
	sync.Mutex
	addr         string
	user         string
	password     string
	db           string
	maxIdleConns int
	idleConns    *list.List
	connCount    int32
}

func Open(addr string, user string, password string, dbName string) (*DB, error) {
	db := &DB{
		addr:      addr,
		user:      user,
		password:  password,
		db:        dbName,
		idleConns: list.New(),
	}

	return db, nil
}

func (db *DB) Addr() string {
	return db.addr
}

func (db *DB) String() string {
	return fmt.Sprintf("%s:%s@%s/%s?maxIdleConns=%v",
		db.user, db.password, db.addr, db.db, db.maxIdleConns)
}

func (db *DB) Close() error {
	db.Lock()
	defer db.Unlock()

	for {
		if db.idleConns.Len() > 0 {
			v := db.idleConns.Back()
			co := v.Value.(*MySqlConn)
			db.idleConns.Remove(v)
			co.Close()
		} else {
			break
		}
	}

	return nil
}

func (db *DB) SetMaxIdleConnNum(num int) {
	db.maxIdleConns = num
}

func (db *DB) GetIdleConnNum() int {
	return db.idleConns.Len()
}

func (db *DB) GetConnNum() int {
	return int(atomic.LoadInt32(&db.connCount))
}

func (db *DB) newConn() (*MySqlConn, error) {
	co := new(MySqlConn)
	if err := co.Connect(db.addr, db.user, db.password, db.db); err != nil {
		return nil, errors.Trace(err)
	}

	return co, nil
}

func (db *DB) PopConn() (co *MySqlConn, err error) {
	db.Lock()
	if db.idleConns.Len() > 0 {
		v := db.idleConns.Front()
		co = v.Value.(*MySqlConn)
		db.idleConns.Remove(v)
	}
	db.Unlock()

	if co != nil {
		//todo: maybe retry
		return co, err
	}

	co, err = db.newConn()
	if err == nil {
		atomic.AddInt32(&db.connCount, 1)
	}

	return co, errors.Trace(err)
}

func (db *DB) PushConn(co *MySqlConn, err error) {
	var closeConn *MySqlConn

	if err != nil {
		closeConn = co
	} else {
		if db.maxIdleConns > 0 {
			db.Lock()

			if db.idleConns.Len() >= db.maxIdleConns {
				v := db.idleConns.Front()
				closeConn = v.Value.(*MySqlConn)
				db.idleConns.Remove(v)
			}

			db.idleConns.PushBack(co)
			db.Unlock()
		} else {
			closeConn = co
		}
	}

	if closeConn != nil {
		atomic.AddInt32(&db.connCount, -1)
		closeConn.Close()
	}
}

type SqlConn struct {
	*MySqlConn
	db *DB
}

func (p *SqlConn) String() string {
	return fmt.Sprintf("{MySqlConn:%+v, db:%+v}", p.MySqlConn, p.db)
}

func (p *SqlConn) Close() {
	if p.MySqlConn != nil {
		p.db.PushConn(p.MySqlConn, p.MySqlConn.pkgErr)
		p.MySqlConn = nil
	}
}

func (db *DB) GetConn() (*SqlConn, error) {
	c, err := db.PopConn()
	return &SqlConn{c, db}, errors.Trace(err)
}
