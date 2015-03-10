package proxy

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/juju/errors"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/hack"
	"github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/sqlparser"
	"github.com/wandoulabs/cm/sqltypes"
	"github.com/wandoulabs/cm/vt/schema"
	"github.com/wandoulabs/cm/vt/tabletserver"
	"github.com/wandoulabs/cm/vt/tabletserver/planbuilder"
)

func applyFilter(columnNumbers []int, input mysql.RowValue) (output mysql.RowValue) {
	output = make(mysql.RowValue, len(columnNumbers))
	for colIndex, colPointer := range columnNumbers {
		if colPointer >= 0 {
			output[colIndex] = input[colPointer]
		}
	}

	return output
}

func (c *Conn) handleQuery(sql string) (err error) {
	sql = strings.TrimRight(sql, ";")
	stmt, err := sqlparser.Parse(sql, c.alloc)
	if err != nil {
		log.Warning(c.connectionId, sql, err)
		return c.handleShow(stmt, sql, nil)
	}

	log.Debugf("connectionId: %d, statement %T , %s", c.connectionId, stmt, sql)

	switch v := stmt.(type) {
	case *sqlparser.Select:
		c.server.IncCounter("select")
		return c.handleSelect(v, sql, nil)
	case *sqlparser.Insert:
		c.server.IncCounter("insert")
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Replace:
		c.server.IncCounter("replace")
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Update:
		c.server.IncCounter("update")
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Delete:
		c.server.IncCounter("delete")
		return c.handleExec(stmt, sql, nil)
	case *sqlparser.Set:
		c.server.IncCounter("set")
		return c.handleSet(v, sql)
	case *sqlparser.SimpleSelect:
		c.server.IncCounter("simple_select")
		return c.handleSimpleSelect(sql, v)
	case *sqlparser.Begin:
		c.server.IncCounter("begin")
		return c.handleBegin()
	case *sqlparser.Commit:
		c.server.IncCounter("commit")
		return c.handleCommit()
	case *sqlparser.Rollback:
		c.server.IncCounter("rollback")
		return c.handleRollback()
	case *sqlparser.Other:
		c.server.IncCounter("other")
		log.Warning(sql)
		return c.handleShow(stmt, sql, nil)
	default:
		return errors.Errorf("statement %T not support now, %+v, %s", stmt, stmt, sql)
	}
}

func (c *Conn) getShardList(stmt sqlparser.Statement) ([]*Shard, error) {
	var shards []*Shard
	ids := c.server.GetShardIds()
	if len(ids) > 0 {
		shards = append(shards, c.server.GetShard(ids[0]))
	}

	//todo: using router info

	return shards, nil
}

func (c *Conn) getConn(n *Shard, isSelect bool) (co *mysql.SqlConn, err error) {
	if !c.needBeginTx() {
		co, err = n.getMasterConn()
		if err != nil {
			return nil, errors.Trace(err)
		}
	} else {
		log.Info("needBeginTx", c.status)
		var ok bool
		co, ok = c.txConns[n.cfg.Id]

		if !ok {
			if co, err = n.getMasterConn(); err != nil {
				return nil, errors.Trace(err)
			}

			log.Debugf("%+v", co)

			if err = co.Begin(); err != nil {
				return nil, errors.Trace(err)
			}

			c.txConns[n.cfg.Id] = co
		}
	}

	if err = co.UseDB(c.db); err != nil {
		return nil, errors.Trace(err)
	}

	if err = co.SetCharset(c.charset); err != nil {
		return nil, errors.Trace(err)
	}

	return
}

func (c *Conn) getShardConns(isSelect bool, stmt sqlparser.Statement) ([]*mysql.SqlConn, error) {
	shards, err := c.getShardList(stmt)
	if err != nil {
		return nil, errors.Trace(err)
	} else if shards == nil {
		return nil, nil
	}

	conns := make([]*mysql.SqlConn, 0, len(shards))

	var co *mysql.SqlConn
	for _, n := range shards {
		co, err = c.getConn(n, isSelect)
		if err != nil {
			log.Error(errors.ErrorStack(err))
			break
		}

		conns = append(conns, co)
	}

	return conns, errors.Trace(err)
}

func (c *Conn) executeInShard(conns []*mysql.SqlConn, sql string, args []interface{}) ([]*mysql.Result, error) {
	wg := &sync.WaitGroup{}
	wg.Add(len(conns))

	rs := make([]interface{}, len(conns))

	for i, co := range conns {
		c.server.AsynExec(
			&execTask{
				wg:   wg,
				rs:   rs,
				idx:  i,
				co:   co,
				sql:  sql,
				args: args,
			})
	}

	wg.Wait()

	var err error
	r := make([]*mysql.Result, len(conns))
	for i, v := range rs {
		if e, ok := v.(error); ok {
			err = e
			break
		}
		r[i] = rs[i].(*mysql.Result)
	}

	return r, errors.Trace(err)
}

func (c *Conn) closeShardConns(conns []*mysql.SqlConn) {
	if c.needBeginTx() {
		return
	}

	for _, co := range conns {
		co.Close()
	}
}

func (c *Conn) newEmptyResultset(stmt *sqlparser.Select) *mysql.Resultset {
	r := &mysql.Resultset{}
	r.Fields = make([]*mysql.Field, len(stmt.SelectExprs))

	for i, expr := range stmt.SelectExprs {
		r.Fields[i] = &mysql.Field{}
		switch e := expr.(type) {
		case *sqlparser.StarExpr:
			r.Fields[i].Name = []byte("*")
		case *sqlparser.NonStarExpr:
			if e.As != nil {
				r.Fields[i].Name = e.As
				r.Fields[i].OrgName = hack.Slice(nstring(e.Expr, c.alloc))
			} else {
				r.Fields[i].Name = hack.Slice(nstring(e.Expr, c.alloc))
			}
		default:
			r.Fields[i].Name = hack.Slice(nstring(e, c.alloc))
		}
	}

	r.Values = make([]mysql.RowValue, 0)
	r.RowDatas = make([]mysql.RowData, 0)

	return r
}

func (c *Conn) getTableSchema(tableName string) (table *schema.Table, ok bool) {
	schemaInfo, ok := c.server.GetRowCacheSchema(c.db)
	if !ok {
		return nil, false
	}

	ti := schemaInfo.GetTable(tableName)
	if ti == nil {
		log.Debug("check if system table", tableName)
		if strings.Index(strings.ToLower(tableName), "information_schema") >= 0 { //system table
			return &schema.Table{
				Name:      tableName,
				CacheType: schema.CACHE_NONE,
			}, true
		} else {
			return nil, false
		}
	}

	log.Infof("%+v", ti.Table)

	return ti.Table, true
}

func (c *Conn) getTableInfo(tableName string) *tabletserver.TableInfo {
	schema, ok := c.server.GetRowCacheSchema(c.db)
	if !ok {
		return nil
	}

	return schema.GetTable(tableName)
}

func (c *Conn) getPlanAndTableInfo(stmt sqlparser.Statement) (*planbuilder.ExecPlan, *tabletserver.TableInfo, error) {
	plan, err := planbuilder.GetStmtExecPlan(stmt, c.getTableSchema, c.alloc)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}

	log.Infof("%+v", plan)

	ti := c.getTableInfo(plan.TableName)

	return plan, ti, nil
}

func pkValuesToStrings(PKColumns []int, pkValues []interface{}) []string {
	composedPkCnt := len(PKColumns)
	s := make([]string, 0, len(pkValues))
	var composedPk string
	for i, values := range pkValues {
		switch v := values.(type) {
		case sqltypes.Value:
			//todo: optimization
			composedPk += v.String()
			composedPk += "--"
			if i%composedPkCnt == composedPkCnt-1 {
				//todo:handle tab
				composedPk = strings.Replace(composedPk, " ", "_", -1)
				s = append(s, composedPk)
				composedPk = "" //reset
			}
		case []interface{}:
			for _, value := range v {
				//todo: optimization
				composedPk += value.(sqltypes.Value).String()
				composedPk += "--"
			}

			if i%composedPkCnt == composedPkCnt-1 {
				//todo:handle tab
				composedPk = strings.Replace(composedPk, " ", "_", -1)
				s = append(s, composedPk)
				composedPk = "" //reset
			}
		default:
			log.Fatal(v, reflect.TypeOf(v))
		}
	}

	return s
}

func getFieldNames(plan *planbuilder.ExecPlan, ti *tabletserver.TableInfo) []schema.TableColumn {
	fields := make([]schema.TableColumn, 0, len(plan.ColumnNumbers)) //construct field name
	for _, i := range plan.ColumnNumbers {
		fields = append(fields, ti.Columns[i])
	}

	return fields
}

//todo: test select a == b && c == d
//select c ==d && a == b
func generateSelectSql(ti *tabletserver.TableInfo, plan *planbuilder.ExecPlan) (string, error) {
	if len(ti.PKColumns) != len(plan.PKValues) {
		log.Error("PKColumns and PKValues not match")
		return "", errors.Errorf("PKColumns and PKValues not match, %+v, %+v", ti.PKColumns, plan.PKValues)
	}

	pks := make([]schema.TableColumn, 0, len(ti.PKColumns))
	for i, _ := range ti.PKColumns {
		pks = append(pks, ti.Columns[ti.PKColumns[i]])
	}

	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("select * from %s where ", ti.Name))
	for i, pk := range pks {
		buf.WriteString(pk.Name)
		buf.WriteString("=")
		plan.PKValues[i].(sqltypes.Value).EncodeSql(buf)
		if i < len(pks)-1 {
			buf.WriteString(" and ")
		}
	}

	buf.WriteString(";")

	return buf.String(), nil
}

func (c *Conn) handleShow(stmt sqlparser.Statement /*Other*/, sql string, args []interface{}) error {
	log.Debug(sql)
	conns, err := c.getShardConns(true, stmt)
	if err != nil {
		return errors.Trace(err)
	} else if len(conns) == 0 {
		return errors.Errorf("not enough connection for %s", sql)
	}

	var rs []*mysql.Result
	rs, err = c.executeInShard(conns, sql, args)
	defer c.closeShardConns(conns)
	if err != nil {
		return errors.Trace(err)
	}

	r := rs[0].Resultset
	status := c.status | rs[0].Status

	log.Debugf("%+v", rs[0])

	//todo: handle set command when sharding
	if stmt == nil {
		log.Warning(sql)
		err := c.writeOkFlush(rs[0])
		return errors.Trace(err)
	}

	for i := 1; i < len(rs); i++ {
		status |= rs[i].Status
		for j := range rs[i].Values {
			r.Values = append(r.Values, rs[i].Values[j])
			r.RowDatas = append(r.RowDatas, rs[i].RowDatas[j])
		}
	}

	return errors.Trace(c.writeResultset(status, r))
}

func (c *Conn) beginShardConns(conns []*mysql.SqlConn) error {
	if c.inTransaction() {
		return nil
	}

	for _, co := range conns {
		if err := co.Begin(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) commitShardConns(conns []*mysql.SqlConn) error {
	if c.inTransaction() {
		return nil
	}

	for _, co := range conns {
		if err := co.Commit(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Conn) mergeExecResult(rs []*mysql.Result) error {
	r := &mysql.Result{}

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

	return errors.Trace(c.writeOK(r))
}

func (c *Conn) mergeSelectResult(rs []*mysql.Result, stmt *sqlparser.Select) error {
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
	/*
		if err := c.limitSelectResult(r, stmt); err != nil {
			return errors.Trace(err)
		}
	*/

	return c.writeResultset(status, r)
}

func (c *Conn) sortSelectResult(r *mysql.Resultset, stmt *sqlparser.Select) error {
	if stmt.OrderBy == nil {
		return nil
	}

	sk := make([]mysql.SortKey, len(stmt.OrderBy))

	for i, o := range stmt.OrderBy {
		sk[i].Name = nstring(o.Expr, c.alloc)
		sk[i].Direction = o.Direction
	}

	return r.Sort(sk)
}

func (c *Conn) limitSelectResult(r *mysql.Resultset, stmt *sqlparser.Select) error {
	if stmt.Limit == nil {
		return nil
	}

	var offset, count int64
	var err error
	if stmt.Limit.Offset == nil {
		offset = 0
	} else {
		if o, ok := stmt.Limit.Offset.(sqlparser.NumVal); !ok {
			return errors.Errorf("invalid select limit %s", nstring(stmt.Limit, c.alloc))
		} else {
			if offset, err = strconv.ParseInt(hack.String([]byte(o)), 10, 64); err != nil {
				return errors.Trace(err)
			}
		}
	}

	if o, ok := stmt.Limit.Rowcount.(sqlparser.NumVal); !ok {
		return errors.Errorf("invalid limit %s", nstring(stmt.Limit, c.alloc))
	} else {
		if count, err = strconv.ParseInt(hack.String([]byte(o)), 10, 64); err != nil {
			return errors.Trace(err)
		} else if count < 0 {
			return errors.Errorf("invalid limit %s", nstring(stmt.Limit, c.alloc))
		}
	}

	if offset+count > int64(len(r.Values)) {
		count = int64(len(r.Values)) - offset
	}

	r.Values = r.Values[offset : offset+count]
	r.RowDatas = r.RowDatas[offset : offset+count]

	return nil
}
