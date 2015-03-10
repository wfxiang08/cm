package proxy

import (
	"github.com/juju/errors"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/hack"
	"github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/sqlparser"
	"github.com/wandoulabs/cm/vt/schema"
	"github.com/wandoulabs/cm/vt/tabletserver"
	"github.com/wandoulabs/cm/vt/tabletserver/planbuilder"
)

type CacheFilter struct{}

func invalidCache(ti *tabletserver.TableInfo, keys []string) {
	for _, key := range keys {
		ti.Cache.Delete(key)
	}
}

func (f *CacheFilter) writeCacheResults(c *Conn, plan *planbuilder.ExecPlan, ti *tabletserver.TableInfo, keys []string, items map[string]tabletserver.RCResult) error {
	values := make([]mysql.RowValue, 0, len(keys))
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

func (f *CacheFilter) fillCacheAndReturnResults(c *Conn, plan *planbuilder.ExecPlan, ti *tabletserver.TableInfo, keys []string) error {
	rowsql, err := generateSelectSql(ti, plan)
	log.Info(rowsql)

	ti.Lock.Lock(hack.Slice(keys[0]))
	defer ti.Lock.Unlock(hack.Slice(keys[0]))

	conns, err := c.getShardConns(true, nil)
	if err != nil {
		return errors.Trace(err)
	} else if len(conns) == 0 {
		return errors.Errorf("not enough connection for %s", rowsql)
	}

	rs, err := c.executeInShard(conns, rowsql, nil)
	defer c.closeShardConns(conns)
	if err != nil {
		return errors.Trace(err)
	}

	//todo:fix hard code
	result := rs[0]

	if len(result.Values) == 0 {
		log.Debug("empty set")
		return c.writeResultset(result.Status, result.Resultset)
	}

	retValues := applyFilter(plan.ColumnNumbers, result.Values[0])

	r, err := c.buildResultset(getFieldNames(plan, ti), []mysql.RowValue{retValues})
	if err != nil {
		log.Error(err)
		return errors.Trace(err)
	}

	//just do simple cache now
	if len(result.Values) == 1 && len(keys) == 1 && ti.CacheType != schema.CACHE_NONE {
		pks := pkValuesToStrings(ti.PKColumns, plan.PKValues)
		log.Debug("fill cache", pks)
		c.server.IncCounter("fill")
		ti.Cache.Set(pks[0], result.RowDatas[0], 0)
	}

	return c.writeResultset(c.status, r)
}

func (f *CacheFilter) OnSelect(c *Conn, stmt *sqlparser.Select, sql string, args []interface{}) (forward bool, err error) {
	// handle cache
	plan, ti, err := c.getPlanAndTableInfo(stmt)
	if err != nil {
		return false, errors.Trace(err)
	}

	log.Debugf("handleSelect %s, %+v", sql, plan.PKValues)

	c.server.IncCounter(plan.PlanId.String())

	if ti != nil && len(plan.PKValues) > 0 && ti.CacheType != schema.CACHE_NONE {
		pks := pkValuesToStrings(ti.PKColumns, plan.PKValues)
		items := ti.Cache.Get(pks, ti.Columns)
		count := 0
		for _, item := range items {
			if item.Row != nil {
				count++
			}
		}

		if count == len(pks) {
			c.server.IncCounter("hint")
			log.Info("hit cache!", sql, pks)
			return false, f.writeCacheResults(c, plan, ti, pks, items)
		}

		c.server.IncCounter("miss")

		if plan.PlanId == planbuilder.PLAN_PK_IN && len(pks) == 1 {
			log.Infof("%s, %+v, %+v", sql, plan, stmt)
			return false, f.fillCacheAndReturnResults(c, plan, ti, pks)
		}
	}

	return true, nil
}

func (f *CacheFilter) OnExec(c *Conn, stmt sqlparser.Statement, sql string, arg []interface{}) (forward bool, err error) {
	// skip cache when insert
	switch stmt.(type) {
	case *sqlparser.Insert:
		return true, nil
	}

	plan, ti, err := c.getPlanAndTableInfo(stmt)
	if err != nil {
		return false, errors.Trace(err)
	}

	if ti == nil {
		return false, errors.Errorf("sql: %s not support", sql)
	}

	c.server.IncCounter(plan.PlanId.String())

	if ti.CacheType != schema.CACHE_NONE {
		if len(ti.PKColumns) != len(plan.PKValues) {
			return false, errors.Errorf("updated/delete/replace without primary key not allowed %+v", plan.PKValues)
		}

		if len(plan.PKValues) == 0 {
			return false, errors.Errorf("pk not exist, sql: %s", sql)
		}

		log.Debugf("%s %+v, %+v", sql, plan, plan.PKValues)
		pks := pkValuesToStrings(ti.PKColumns, plan.PKValues)

		ti.Lock.Lock(hack.Slice(pks[0]))
		defer ti.Lock.Unlock(hack.Slice(pks[0]))

		invalidCache(ti, pks)
	}

	return true, nil
}
