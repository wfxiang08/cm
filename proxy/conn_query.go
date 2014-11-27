package proxy

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/juju/errors"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/hack"
	. "github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/sqlparser"
	"github.com/wandoulabs/cm/sqltypes"
	"github.com/wandoulabs/cm/vt/schema"
	"github.com/wandoulabs/cm/vt/tabletserver/planbuilder"
)

func applyFilter(columnNumbers []int, input RowValue) (output RowValue) {
	output = make(RowValue, len(columnNumbers))
	for colIndex, colPointer := range columnNumbers {
		if colPointer >= 0 {
			output[colIndex] = input[colPointer]
		}
	}

	return output
}

func (c *Conn) handleQuery(sql string) (err error) {
	sql = strings.TrimRight(sql, ";")
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return errors.Errorf(`parse sql "%s" error`, sql)
	}

	switch v := stmt.(type) {
	case *sqlparser.Select:
		return c.handleSelect(v, sql, nil)
	case *sqlparser.Insert:
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Update:
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Delete:
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Set:
		return c.handleSet(v)
	case *sqlparser.SimpleSelect:
		return c.handleSimpleSelect(sql, v)
	default:
		return errors.Errorf("statement %T not support now, %+v, %s", stmt, stmt, sql)
	}

	return nil
}

func (c *Conn) getShardList(stmt sqlparser.Statement, bindVars map[string]interface{}) ([]*Node, error) {
	if c.schema == nil {
		return nil, NewDefaultError(ER_NO_DB_ERROR)
	}

	var n []*Node
	names := c.server.getNodeNames()
	if len(names) > 0 {
		n = append(n, c.server.getNode(names[0]))
	}
	return n, nil
}

func (c *Conn) getConn(n *Node, isSelect bool) (co *SqlConn, err error) {
	if !c.needBeginTx() {
		if isSelect {
			co, err = n.getSelectConn()
		} else {
			co, err = n.getMasterConn()
		}
		if err != nil {
			return
		}
	} else {
		var ok bool
		c.Lock()
		co, ok = c.txConns[n]
		c.Unlock()

		if !ok {
			if co, err = n.getMasterConn(); err != nil {
				return
			}

			if err = co.Begin(); err != nil {
				return
			}

			c.Lock()
			c.txConns[n] = co
			c.Unlock()
		}
	}

	if err = co.UseDB(c.schema.db); err != nil {
		return
	}

	if err = co.SetCharset(c.charset); err != nil {
		return
	}

	return
}

func (c *Conn) getShardConns(isSelect bool, stmt sqlparser.Statement, bindVars map[string]interface{}) ([]*SqlConn, error) {
	nodes, err := c.getShardList(stmt, bindVars)
	if err != nil {
		return nil, err
	} else if nodes == nil {
		return nil, nil
	}

	conns := make([]*SqlConn, 0, len(nodes))

	var co *SqlConn
	for _, n := range nodes {
		co, err = c.getConn(n, isSelect)
		if err != nil {
			log.Info(err)
			break
		}

		conns = append(conns, co)
	}

	return conns, errors.Trace(err)
}

func (c *Conn) executeInShard(conns []*SqlConn, sql string, args []interface{}) ([]*Result, error) {
	var wg sync.WaitGroup
	wg.Add(len(conns))

	rs := make([]interface{}, len(conns))
	f := func(rs []interface{}, i int, co *SqlConn) {
		r, err := co.Execute(sql, args...)
		if err != nil {
			rs[i] = err
		} else {
			rs[i] = r
		}

		wg.Done()
	}

	for i, co := range conns {
		go f(rs, i, co)
	}

	wg.Wait()

	var err error
	r := make([]*Result, len(conns))
	for i, v := range rs {
		if e, ok := v.(error); ok {
			err = e
			break
		}
		r[i] = rs[i].(*Result)
	}

	return r, errors.Trace(err)
}

func (c *Conn) closeShardConns(conns []*SqlConn, rollback bool) {
	if c.isInTransaction() {
		return
	}

	for _, co := range conns {
		if rollback {
			co.Rollback()
		}

		co.Close()
	}
}

func (c *Conn) newEmptyResultset(stmt *sqlparser.Select) *Resultset {
	r := new(Resultset)
	r.Fields = make([]*Field, len(stmt.SelectExprs))

	for i, expr := range stmt.SelectExprs {
		r.Fields[i] = &Field{}
		switch e := expr.(type) {
		case *sqlparser.StarExpr:
			r.Fields[i].Name = []byte("*")
		case *sqlparser.NonStarExpr:
			if e.As != nil {
				r.Fields[i].Name = e.As
				r.Fields[i].OrgName = hack.Slice(nstring(e.Expr))
			} else {
				r.Fields[i].Name = hack.Slice(nstring(e.Expr))
			}
		default:
			r.Fields[i].Name = hack.Slice(nstring(e))
		}
	}

	r.Values = make([]RowValue, 0)
	r.RowDatas = make([]RowData, 0)

	return r
}

func makeBindVars(args []interface{}) map[string]interface{} {
	bindVars := make(map[string]interface{}, len(args))

	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i+1)] = v
	}

	return bindVars
}

func (c *Conn) handleSelect(stmt *sqlparser.Select, sql string, args []interface{}) error {
	// handle cache
	GetTable := func(tableName string) (table *schema.Table, ok bool) {
		ti := c.server.autoSchamas[c.db].GetTable(tableName)
		if ti == nil {
			return nil, false
		}

		log.Infof("%+v", ti.Table)

		return ti.Table, true
	}

	plan, err := planbuilder.GetExecPlan(sql, GetTable)
	if err != nil {
		return errors.Trace(err)
	}

	//todo: fix hard code
	ti := c.server.autoSchamas[c.db].GetTable(plan.TableName)
	if ti == nil {
		return errors.Errorf("unsupport sql %s", sql)
	}

	key := plan.PKValues[0].(sqltypes.Value).String()
	items := ti.Cache.Get([]string{key})
	if row, ok := items[key]; ok {
		retValue := applyFilter(plan.ColumnNumbers, row.Row)
		var fields []string
		var values []RowValue
		for _, i := range plan.ColumnNumbers {
			c := ti.Columns[i]
			fields = append(fields, c.Name)
		}
		values = append(values, retValue)
		r, err := c.buildResultset(fields, values)
		if err != nil {
			log.Error(err)
		}
		//todo:write back
		log.Info("hit cache!")
		return c.writeResultset(c.status, r)
	}

	pk := ti.Columns[ti.PKColumns[0]]
	rowsql := fmt.Sprintf("select * from %s where %s = %s", plan.TableName, pk.Name, plan.PKValues[0])
	log.Info(rowsql)

	result, err := c.server.autoSchamas[c.db].Exec(rowsql)
	if err != nil {
		return errors.Trace(err)
	}

	log.Infof("%s, %+v, %+v", sql, plan, stmt)
	retValues := applyFilter(plan.ColumnNumbers, result.Values[0])
	log.Infof("%+v", retValues)
	ti.Cache.Set(key, result.Values[0], 0)

	bindVars := makeBindVars(args)

	conns, err := c.getShardConns(true, stmt, bindVars)
	if err != nil {
		return errors.Trace(err)
	} else if conns == nil {
		r := c.newEmptyResultset(stmt)
		return c.writeResultset(c.status, r)
	}

	var rs []*Result
	rs, err = c.executeInShard(conns, sql, args)
	c.closeShardConns(conns, false)
	if err == nil {
		err = c.mergeSelectResult(rs, stmt)
	}

	return errors.Trace(err)
}

func (c *Conn) handleExec(stmt sqlparser.Statement, sql string, args []interface{}) error {
	bindVars := makeBindVars(args)

	conns, err := c.getShardConns(false, stmt, bindVars)
	if err != nil {
		return errors.Trace(err)
	} else if conns == nil {
		return c.writeOK(nil)
	}

	var rs []*Result
	if len(conns) == 1 {
		rs, err = c.executeInShard(conns, sql, args)
	} else {
		log.Warning("not implement yet")

	}

	c.closeShardConns(conns, err != nil)

	if err == nil {
		err = c.mergeExecResult(rs)
	}

	return errors.Trace(err)
}

func (c *Conn) mergeExecResult(rs []*Result) error {
	r := new(Result)

	for _, v := range rs {
		r.Status |= v.Status
		r.AffectedRows += v.AffectedRows
		if r.InsertId == 0 {
			r.InsertId = v.InsertId
		} else if r.InsertId > v.InsertId {

			r.InsertId = v.InsertId
		}
	}

	if r.InsertId > 0 {
		c.lastInsertId = int64(r.InsertId)
	}

	c.affectedRows = int64(r.AffectedRows)

	return c.writeOK(r)
}

func (c *Conn) mergeSelectResult(rs []*Result, stmt *sqlparser.Select) error {
	r := rs[0].Resultset

	status := c.status | rs[0].Status

	for i := 1; i < len(rs); i++ {
		status |= rs[i].Status

		for j := range rs[i].Values {
			r.Values = append(r.Values, rs[i].Values[j])
			r.RowDatas = append(r.RowDatas, rs[i].RowDatas[j])
		}
	}

	c.sortSelectResult(r, stmt)

	if err := c.limitSelectResult(r, stmt); err != nil {
		return errors.Trace(err)
	}

	return c.writeResultset(status, r)
}

func (c *Conn) sortSelectResult(r *Resultset, stmt *sqlparser.Select) error {
	if stmt.OrderBy == nil {
		return nil
	}

	sk := make([]SortKey, len(stmt.OrderBy))

	for i, o := range stmt.OrderBy {
		sk[i].Name = nstring(o.Expr)
		sk[i].Direction = o.Direction
	}

	return r.Sort(sk)
}

func (c *Conn) limitSelectResult(r *Resultset, stmt *sqlparser.Select) error {
	if stmt.Limit == nil {
		return nil
	}

	var offset, count int64
	var err error
	if stmt.Limit.Offset == nil {
		offset = 0
	} else {
		if o, ok := stmt.Limit.Offset.(sqlparser.NumVal); !ok {
			return errors.Errorf("invalid select limit %s", nstring(stmt.Limit))
		} else {
			if offset, err = strconv.ParseInt(hack.String([]byte(o)), 10, 64); err != nil {
				return errors.Trace(err)
			}
		}
	}

	if o, ok := stmt.Limit.Rowcount.(sqlparser.NumVal); !ok {
		return errors.Errorf("invalid limit %s", nstring(stmt.Limit))
	} else {
		if count, err = strconv.ParseInt(hack.String([]byte(o)), 10, 64); err != nil {
			return errors.Trace(err)
		} else if count < 0 {
			return errors.Errorf("invalid limit %s", nstring(stmt.Limit))
		}
	}

	if offset+count > int64(len(r.Values)) {
		count = int64(len(r.Values)) - offset
	}

	r.Values = r.Values[offset : offset+count]
	r.RowDatas = r.RowDatas[offset : offset+count]

	return nil
}
