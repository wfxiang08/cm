package proxy

import (
	"github.com/juju/errors"
	"github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/sqlparser"
)

func (c *Conn) handleExec(stmt sqlparser.Statement, sql string, args []interface{}, skipCache bool) error {
	conns, err := c.getShardConns(false, stmt)
	if err != nil {
		return errors.Trace(err)
	} else if len(conns) == 0 { //todo:handle error
		return errors.Errorf("not server found %s", sql)
	}

	var rs []*mysql.Result
	rs, err = c.executeInShard(conns, sql, args)

	c.closeShardConns(conns)

	if err == nil {
		err = c.mergeExecResult(rs)
	}

	return errors.Trace(err)
}
