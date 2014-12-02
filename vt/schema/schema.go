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

func (ta *Table) AddColumn(name string, columnType string, defval mysql.Value, extra string) {
	index := len(ta.Columns)
	ta.Columns = append(ta.Columns, TableColumn{Name: name})
	if strings.Contains(columnType, "int") || strings.Contains(columnType, "long") || strings.Contains(columnType, "tiny") || strings.Contains(columnType, "short") {
		ta.Columns[index].Category = mysql.MYSQL_TYPE_LONGLONG
	} else if strings.HasPrefix(columnType, "varbinary") || strings.Contains(columnType, "blob") {
		ta.Columns[index].Category = mysql.MYSQL_TYPE_VARCHAR
	} else if strings.HasPrefix(columnType, "datetime") || strings.HasPrefix(columnType, "timestamp") {
		ta.Columns[index].Category = mysql.MYSQL_TYPE_DATETIME
	} else if strings.HasPrefix(columnType, "date") {
		ta.Columns[index].Category = mysql.MYSQL_TYPE_DATE
	} else if strings.Contains(columnType, "float") || strings.Contains(columnType, "double") {
		ta.Columns[index].Category = mysql.MYSQL_TYPE_DOUBLE
	} else if strings.HasPrefix(columnType, "enum") {
		ta.Columns[index].Category = mysql.MYSQL_TYPE_ENUM
	} else if strings.Contains(columnType, "text") || strings.Contains(columnType, "varchar") || strings.Contains(columnType, "string") || strings.Contains(columnType, "char") {
		ta.Columns[index].Category = mysql.MYSQL_TYPE_STRING
	} else {
		log.Fatalf("not support type: %s", columnType)
	}
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
