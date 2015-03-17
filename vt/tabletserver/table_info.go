// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabletserver

import (
	"fmt"
	"strings"

	"github.com/juju/errors"
	"github.com/ngaut/lockring"
	log "github.com/ngaut/logging"
	"github.com/ngaut/sync2"
	"github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/sqltypes"
	"github.com/wandoulabs/cm/vt/schema"
)

type TableInfo struct {
	Lock *lockring.LockRing
	*schema.Table
	// stats updated by sqlquery.go
	hits, absent, misses, invalidations sync2.AtomicInt64
}

func NewTableInfo(conn *mysql.MySqlConn, tableName string, tableType string, createTime sqltypes.Value,
	comment string) (ti *TableInfo, err error) {
	ti, err = loadTableInfo(conn, tableName)
	if err != nil {
		return nil, errors.Trace(err)
	}

	ti.Lock = lockring.New(65536)

	return ti, nil
}

func loadTableInfo(conn *mysql.MySqlConn, tableName string) (ti *TableInfo, err error) {
	ti = &TableInfo{Table: schema.NewTable(tableName)}
	if err = ti.fetchColumns(conn); err != nil {
		return nil, errors.Trace(err)
	}

	if err = ti.fetchIndexes(conn); err != nil {
		return nil, errors.Trace(err)
	}

	return ti, nil
}

func (ti *TableInfo) fetchColumns(conn *mysql.MySqlConn) error {
	columns, err := conn.Execute(fmt.Sprintf("show full columns from `%s`", ti.Name))
	if err != nil {
		return errors.Trace(err)
	}

	for _, row := range columns.Values {
		v, err := sqltypes.BuildValue(row[5])
		if err != nil {
			return errors.Trace(err)
		}

		var collation string
		if row[2] != nil {
			collation = string(row[2].([]byte))
		}
		extra := string(row[6].([]byte))
		columnType := string(row[1].([]byte))
		columnName := string(row[0].([]byte))
		ti.AddColumn(columnName, columnType, collation,
			v, extra)
	}

	log.Debugf("%s %+v", ti.Name, ti.Columns)

	return nil
}

func (ti *TableInfo) SetPK(colnames []string) error {
	log.Debugf("table %s SetPK %s", ti.Name, colnames)
	pkIndex := schema.NewIndex("PRIMARY")
	colnums := make([]int, len(colnames))
	for i, colname := range colnames {
		colnums[i] = ti.FindColumn(strings.ToLower(colname))
		if colnums[i] == -1 {
			return errors.Errorf("column %s not found, %+v", colname, ti.Columns)
		}
		pkIndex.AddColumn(strings.ToLower(colname), 1)
	}

	for _, col := range ti.Columns {
		pkIndex.DataColumns = append(pkIndex.DataColumns, strings.ToLower(col.Name))
	}

	if len(ti.Indexes) == 0 {
		ti.Indexes = make([]*schema.Index, 1)
	} else if ti.Indexes[0].Name != "PRIMARY" {
		ti.Indexes = append(ti.Indexes, nil)
		copy(ti.Indexes[1:], ti.Indexes[:len(ti.Indexes)-1])
	} // else we replace the currunt primary key

	ti.Indexes[0] = pkIndex
	ti.PKColumns = colnums
	return nil
}

func (ti *TableInfo) fetchIndexes(conn *mysql.MySqlConn) error {
	/*
		indexes, err := conn.Execute(fmt.Sprintf("show index from `%s`", ti.Name))
		if err != nil {
			log.Error(err)
			return errors.Trace(err)
		}

		log.Debugf("%+v", indexes.Values)

		var currentIndex *schema.Index
		currentName := ""
		for _, row := range indexes.Values {
			indexName, err := mysql.Raw(row[2])
			if err != nil {
				log.Error(err)
				return errors.Trace(err)
			}

			if currentName != string(indexName) {
				currentIndex = ti.AddIndex(string(indexName))
				currentName = string(indexName)
			}

			var cardinality string
			if row[6] != nil {
				v, err := mysql.Raw(row[6])
				if err != nil {
					log.Warningf("%s", err)
					return errors.Trace(err)
				}
				cardinality = string(v)

			}
			val, err := strconv.ParseUint(cardinality, 0, 64)
			if err != nil {
				log.Warningf("%s", err)
				return errors.Trace(err)
			}
			currentIndex.AddColumn(string(row[4].([]byte)), val)
		}

		log.Debugf("table: %s, indexes: %+v", ti.Name, ti.Indexes)

		if len(ti.Indexes) == 0 {
			return nil
		}

		pkIndex := ti.Indexes[0]
		if pkIndex.Name != "PRIMARY" {
			return nil
		}
		ti.PKColumns = make([]int, len(pkIndex.Columns))
		for i, pkCol := range pkIndex.Columns {
			ti.PKColumns[i] = ti.FindColumn(pkCol)
		}
		// Primary key contains all table columns
		for _, col := range ti.Columns {
			pkIndex.DataColumns = append(pkIndex.DataColumns, col.Name)
		}
		// Secondary indices contain all primary key columns
		for i := 1; i < len(ti.Indexes); i++ {
			for _, c := range ti.Indexes[i].Columns {
				ti.Indexes[i].DataColumns = append(ti.Indexes[i].DataColumns, c)
			}
			for _, c := range pkIndex.Columns {
				// pk columns may already be part of the index. So,
				// check before adding.
				if ti.Indexes[i].FindDataColumn(c) != -1 {
					continue
				}
				ti.Indexes[i].DataColumns = append(ti.Indexes[i].DataColumns, c)
			}
		}
	*/

	return nil
}

func (ti *TableInfo) StatsJSON() string {
	h, a, m, i := ti.Stats()
	return fmt.Sprintf("{\"Hits\": %v, \"Absent\": %v, \"Misses\": %v, \"Invalidations\": %v}", h, a, m, i)
}

func (ti *TableInfo) Stats() (hits, absent, misses, invalidations int64) {
	return ti.hits.Get(), ti.absent.Get(), ti.misses.Get(), ti.invalidations.Get()
}
