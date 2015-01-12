package proxy

import (
	"strings"

	"github.com/juju/errors"
	. "github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/sqlparser"
)

var nstring = sqlparser.String

func (c *Conn) handleSet(stmt *sqlparser.Set, sql string) error {
	if len(stmt.Exprs) != 1 {
		return errors.Errorf("must set one item once, not %s", nstring(stmt, c.alloc))
	}

	k := string(stmt.Exprs[0].Name.Name)

	switch strings.ToUpper(k) {
	case `AUTOCOMMIT`:
		return c.handleSetAutoCommit(stmt.Exprs[0].Expr)
	case `NAMES`:
		return c.handleSetNames(stmt.Exprs[0].Expr)
	default:
		return c.handleShow(nil, sql, nil) //errors.Errorf("set %s is not supported now", k)
	}
}

func (c *Conn) handleSetAutoCommit(val sqlparser.ValExpr) error {
	value, ok := val.(sqlparser.NumVal)
	if !ok {
		return errors.Errorf("set autocommit error")
	}
	switch value[0] {
	case '1':
		c.status |= SERVER_STATUS_AUTOCOMMIT
	case '0':
		c.status &= ^SERVER_STATUS_AUTOCOMMIT
	default:
		return errors.Errorf("invalid autocommit flag %s", value)
	}

	err := c.writeOK(nil)
	if err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(c.flush())
}

func (c *Conn) handleSetNames(val sqlparser.ValExpr) error {
	value, ok := val.(sqlparser.StrVal)
	if !ok {
		return errors.Errorf("set names charset error")
	}

	charset := strings.ToLower(string(value))
	cid, ok := CharsetIds[charset]
	if !ok {
		return errors.Errorf("invalid charset %s", charset)
	}

	c.charset = charset
	c.collation = cid

	err := c.writeOK(nil)
	if err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(c.flush())
}
