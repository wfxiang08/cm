package proxy

import (
	. "github.com/wandoulabs/cm/mysql"
)

func (c *Conn) isInTransaction() bool {
	return c.status&SERVER_STATUS_IN_TRANS > 0
}

func (c *Conn) isAutoCommit() bool {
	return c.status&SERVER_STATUS_AUTOCOMMIT > 0
}

//if status is in_trans, need
//else if status is not autocommit, need
//else no need
func (c *Conn) needBeginTx() bool {
	return c.isInTransaction() || !c.isAutoCommit()
}
