package proxy

import (
	log "github.com/ngaut/logging"
	. "github.com/wandoulabs/cm/mysql"
)

func (c *Conn) isInTransaction() bool {
	return c.status&SERVER_STATUS_IN_TRANS > 0
}

func (c *Conn) isAutoCommit() bool {
	return c.status&SERVER_STATUS_AUTOCOMMIT > 0
}

func (c *Conn) handleBegin() error {
	log.Debug("handle begin")
	c.status |= SERVER_STATUS_IN_TRANS

	return c.writeOkFlush(nil)
}

func (c *Conn) handleCommit() (err error) {
	if err := c.commit(); err != nil {
		return err
	}

	return c.writeOkFlush(nil)
}

func (c *Conn) handleRollback() (err error) {
	if err := c.rollback(); err != nil {
		return err
	}

	return c.writeOkFlush(nil)
}

func (c *Conn) commit() (err error) {
	log.Debugf("handle  commit on %v", c)
	c.status &= ^SERVER_STATUS_IN_TRANS

	for _, co := range c.txConns {
		if e := co.Commit(); e != nil {
			err = e
		}
		co.Close()
	}

	c.txConns = make(map[string]*SqlConn)

	return
}

func (c *Conn) rollback() (err error) {
	log.Debugf("handle  rollback on %v", c)
	c.status &= ^SERVER_STATUS_IN_TRANS

	for _, co := range c.txConns {
		if e := co.Rollback(); e != nil {
			err = e
		}
		co.Close()
	}

	c.txConns = make(map[string]*SqlConn)

	return
}

//if status is in_trans, need
//else if status is not autocommit, need
//else no need
//todo: rename this function
func (c *Conn) needBeginTx() bool {
	return c.isInTransaction() || !c.isAutoCommit()
}
