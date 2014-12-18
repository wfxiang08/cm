// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package schema

// Yes, this sucks. It's a tiny tiny package that needs to be on its own
// It contains a data structure that's shared between sqlparser & tabletserver

import (
	"strings"

	log "github.com/ngaut/logging"

	"github.com/wandoulabs/cm/mysql"
)

// Cache types
const (
	CACHE_NONE = 0
	CACHE_RW   = 1
	CACHE_W    = 2
)

type TableColumn struct {
	Name     string
	Category byte
	IsAuto   bool
	Default  mysql.Value
}

type Table struct {
	Name      string
	Columns   []TableColumn
	Indexes   []*Index
	PKColumns []int
	CacheType int
}

func NewTable(name string) *Table {
	return &Table{
		Name:    name,
		Columns: make([]TableColumn, 0, 16),
		Indexes: make([]*Index, 0, 8),
	}
}

var typesMap = map[string]byte{
	"int":       mysql.MYSQL_TYPE_LONG,
	"long":      mysql.MYSQL_TYPE_LONG,
	"short":     mysql.MYSQL_TYPE_SHORT,
	"tiny":      mysql.MYSQL_TYPE_TINY,
	"varbinary": mysql.MYSQL_TYPE_VARCHAR,
	"blob":      mysql.MYSQL_TYPE_BLOB,
	"datetime":  mysql.MYSQL_TYPE_DATETIME,
	"date":      mysql.MYSQL_TYPE_DATE,
	"timestamp": mysql.MYSQL_TYPE_TIMESTAMP,
	"data":      mysql.MYSQL_TYPE_DATE,
	"float":     mysql.MYSQL_TYPE_FLOAT,
	"double":    mysql.MYSQL_TYPE_DOUBLE,
	"enum":      mysql.MYSQL_TYPE_ENUM,
	"text":      mysql.MYSQL_TYPE_STRING,
	"varchar":   mysql.MYSQL_TYPE_VARCHAR,
	"string":    mysql.MYSQL_TYPE_STRING,
	"char":      mysql.MYSQL_TYPE_STRING,
}

func str2mysqlType(columnType string) byte {
	b, ok := typesMap[columnType]
	if !ok {
		log.Fatalf("%s not exist", columnType)
	}

	return b
}

func (ta *Table) AddColumn(name string, columnType string, defval mysql.Value, extra string) {
	index := len(ta.Columns)
	ta.Columns = append(ta.Columns, TableColumn{Name: name})

	endPos := strings.Index(columnType, "(") //handle something like: int(11)
	if endPos > 0 {
		columnType = strings.ToLower(strings.TrimSpace(columnType[:endPos]))
	}

	ta.Columns[index].Category = str2mysqlType(columnType)

	if extra == "auto_increment" {
		ta.Columns[index].IsAuto = true
		// Ignore default value, if any
		return
	}
	if defval == nil {
		return
	}
	ta.Columns[index].Default = defval
}

func (ta *Table) FindColumn(name string) int {
	for i, col := range ta.Columns {
		if col.Name == name {
			return i
		}
	}

	return -1
}

func (ta *Table) GetPKColumn(index int) *TableColumn {
	return &ta.Columns[ta.PKColumns[index]]
}

func (ta *Table) AddIndex(name string) (index *Index) {
	index = NewIndex(name)
	ta.Indexes = append(ta.Indexes, index)

	return index
}

type Index struct {
	Name        string
	Columns     []string
	Cardinality []uint64
	DataColumns []string
}

func NewIndex(name string) *Index {
	return &Index{name, make([]string, 0, 8), make([]uint64, 0, 8), nil}
}

func (idx *Index) AddColumn(name string, cardinality uint64) {
	idx.Columns = append(idx.Columns, name)
	if cardinality == 0 {
		cardinality = uint64(len(idx.Cardinality) + 1)
	}
	idx.Cardinality = append(idx.Cardinality, cardinality)
}

func (idx *Index) FindColumn(name string) int {
	for i, colName := range idx.Columns {
		if name == colName {
			return i
		}
	}

	return -1
}

func (idx *Index) FindDataColumn(name string) int {
	for i, colName := range idx.DataColumns {
		if name == colName {
			return i
		}
	}

	return -1
}
