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
	. "github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/sqlparser"
	"github.com/wandoulabs/cm/sqltypes"
	"github.com/wandoulabs/cm/vt/schema"
	"github.com/wandoulabs/cm/vt/tabletserver"
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
	log.Debug(sql)
	sql = strings.TrimRight(sql, ";")
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		log.Warning(sql)
		return c.handleShow(stmt, sql, nil)
		//return errors.Errorf(`parse sql "%s" error: %s`, sql, err)
	}

	log.Debugf("statement %T , %+v, %s", stmt, stmt, sql)

	switch v := stmt.(type) {
	case *sqlparser.Select:
		return c.handleSelect(v, sql, nil)
	case *sqlparser.Insert:
		return c.handleExec(stmt, sql, nil, true)
	case *sqlparser.Update:
		return c.handleExec(stmt, sql, nil, false)
	case *sqlparser.Delete:
		return c.handleExec(stmt, sql, nil, false)
	case *sqlparser.Set:
		return c.handleSet(v, sql)
	case *sqlparser.SimpleSelect:
		return c.handleSimpleSelect(sql, v)
	case *sqlparser.Other:
		return c.handleShow(stmt, sql, nil)
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

func (c *Conn) getTableSchema(tableName string) (table *schema.Table, ok bool) {
	ti := c.server.autoSchamas[c.db].GetTable(tableName)
	if ti == nil {
		return nil, false
	}

	log.Infof("%+v", ti.Table)

	return ti.Table, true
}

func (c *Conn) getTableInfo(tableName string) *tabletserver.TableInfo {
	return c.server.autoSchamas[c.db].GetTable(tableName)
}

func (c *Conn) getPlanAndTableInfo(sql string) (*planbuilder.ExecPlan, *tabletserver.TableInfo, error) {
	plan, err := planbuilder.GetExecPlan(sql, c.getTableSchema)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}

	log.Infof("%+v", plan)

	ti := c.getTableInfo(plan.TableName)
	if ti == nil {
		return plan, nil, errors.Errorf("unsupport sql %s", sql)
	}

	return plan, ti, nil
}

func pkValuesToStrings(pkValues []interface{}) []string {
	s := make([]string, 0)
	for _, values := range pkValues {
		switch v := values.(type) {
		case sqltypes.Value:
			s = append(s, v.String())
		case []interface{}:
			for _, value := range v {
				s = append(s, value.(sqltypes.Value).String())
			}
		default:
			log.Fatal(v, reflect.TypeOf(v))
		}
	}

	return s
}

func getFieldNames(plan *planbuilder.ExecPlan, ti *tabletserver.TableInfo) []string {
	fields := make([]string, 0, len(plan.ColumnNumbers)) //construct field name
	for _, i := range plan.ColumnNumbers {
		c := ti.Columns[i]
		fields = append(fields, c.Name)
	}

	return fields
}

func (c *Conn) writeCacheResults(plan *planbuilder.ExecPlan, ti *tabletserver.TableInfo, keys []string, items map[string]tabletserver.RCResult) error {
	var values []RowValue
	for _, key := range keys {
		row, ok := items[key]
		if !ok {
			log.Fatal("should never happend")
		}
		retValue := applyFilter(plan.ColumnNumbers, row.Row)
		values = append(values, retValue)
	}

	r, err := c.buildResultset(getFieldNames(plan, ti), values)
	if err != nil {
		log.Error(err)
		return errors.Trace(err)
	}

	return errors.Trace(c.writeResultset(c.status, r))
}

//todo: test select a == b && c == d
//select c ==d && a == b
func generateSelectSql(ti *tabletserver.TableInfo, plan *planbuilder.ExecPlan) (string, error) {
	if len(ti.PKColumns) != len(plan.PKValues) {
		log.Error("PKColumns and PKValues not match")
		return "", errors.Errorf("PKColumns and PKValues not match", ti.PKColumns, plan.PKValues)
	}

	pks := make([]schema.TableColumn, 0, len(ti.PKColumns))
	for i, _ := range ti.PKColumns {
		pks = append(pks, ti.Columns[ti.PKColumns[i]])
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("select * from %s where ", ti.Name))
	for i, pk := range pks {
		buf.WriteString(pk.Name)
		buf.WriteString("=")
		buf.WriteString(fmt.Sprintf("%s", plan.PKValues[i]))
		if i < len(pks)-1 {
			buf.WriteString(" and ")
		}
	}

	buf.WriteString(";")

	return buf.String(), nil
}

func (c *Conn) fillCacheAndReturnResults(plan *planbuilder.ExecPlan, ti *tabletserver.TableInfo, keys []string) error {
	rowsql, err := generateSelectSql(ti, plan)
	log.Info(rowsql)

	result, err := c.server.autoSchamas[c.db].Exec(rowsql)
	if err != nil {
		return errors.Trace(err)
	}

	if len(result.Values) == 0 {
		log.Debug("empty set")
		return c.writeResultset(result.Status, result.Resultset)
	}

	retValues := applyFilter(plan.ColumnNumbers, result.Values[0])
	//log.Debug(len(retValues), len(keys))

	//just do simple cache now
	if len(result.Values) == 1 && len(keys) == 1 && ti.CacheType != schema.CACHE_NONE {
		pkvalue := plan.PKValues[0].(sqltypes.Value).String()
		log.Debug("fill cache", pkvalue)
		ti.Cache.Set(pkvalue, result.Values[0], 0)
	}

	var values []RowValue
	values = append(values, retValues)
	r, err := c.buildResultset(getFieldNames(plan, ti), values)
	if err != nil {
		log.Error(err)
		return errors.Trace(err)
	}

	return c.writeResultset(c.status, r)
}

func (c *Conn) handleShow(stmt sqlparser.Statement /*Other*/, sql string, args []interface{}) error {
	log.Debug(sql)
	bindVars := makeBindVars(args)
	conns, err := c.getShardConns(true, stmt, bindVars)
	if err != nil {
		return errors.Trace(err)
	} else if conns == nil {
		return errors.Errorf("not enough connection for %s", sql)
	}

	var rs []*Result
	rs, err = c.executeInShard(conns, sql, args)
	c.closeShardConns(conns, false)

	r := rs[0].Resultset
	status := c.status | rs[0].Status

	//todo: handle set command when sharding
	if stmt == nil { //hack for "set names utf8"
		return errors.Trace(c.writeOK(rs[0]))
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

func (c *Conn) handleSelect(stmt *sqlparser.Select, sql string, args []interface{}) error {
	log.Debug("handleSelect", sql)
	// handle cache
	plan, ti, err := c.getPlanAndTableInfo(sql)
	if err != nil {
		return errors.Trace(err)
	}

	if len(plan.PKValues) > 0 && ti.CacheType != schema.CACHE_NONE {
		keys := pkValuesToStrings(plan.PKValues)
		items := ti.Cache.Get(keys)
		count := 0
		for _, item := range items {
			if item.Row != nil {
				count++
			}
		}

		if count == len(keys) { //all cache hint
			log.Info("hit cache!", sql, items, keys)
			return c.writeCacheResults(plan, ti, keys, items)
		}

		log.Debug(count, len(keys), keys)

		if plan.PlanId == planbuilder.PLAN_PK_IN && len(keys) == 1 {
			log.Infof("%s, %+v, %+v", sql, plan, stmt)
			return c.fillCacheAndReturnResults(plan, ti, keys)
		}
	}

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

func invalidCache(ti *tabletserver.TableInfo, keys []string) {
	for _, key := range keys {
		ti.Cache.Delete(key)
	}
}

func (c *Conn) handleExec(stmt sqlparser.Statement, sql string, args []interface{}, skipCache bool) error {
	if !skipCache {
		// handle cache
		plan, ti, err := c.getPlanAndTableInfo(sql)
		if err != nil {
			return errors.Trace(err)
		}

		if ti.CacheType != schema.CACHE_NONE {
			if len(plan.PKValues) == 0 {
				return errors.Errorf("pk not exist, sql: %s", sql)
			}

			log.Debugf("%s %+v, %+v", sql, plan, plan.PKValues)
			//todo: test composed pk
			keys := pkValuesToStrings(plan.PKValues)
			invalidCache(ti, keys)
		}
	}

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
