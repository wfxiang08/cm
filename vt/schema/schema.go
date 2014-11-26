// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package schema

// Yes, this sucks. It's a tiny tiny package that needs to be on its own
// It contains a data structure that's shared between sqlparser & tabletserver

import (
	"strings"

	"github.com/wandoulabs/cm/mysql"
)

// Column categories
const (
	CAT_OTHER = iota
	CAT_INT64
	CAT_FLOAT64
	CAT_BYTES
	CAT_STR
	CAT_TIME
)

// Cache types
const (
	CACHE_NONE = 0
	CACHE_RW   = 1
	CACHE_W    = 2
)

type TableColumn struct {
	Name     string
	Category int
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
	if strings.Contains(columnType, "int") {
		ta.Columns[index].Category = CAT_INT64
	} else if strings.HasPrefix(columnType, "varbinary") {
		ta.Columns[index].Category = CAT_BYTES
	} else if strings.Contains(columnType, "float") || strings.Contains(columnType, "double") {
		ta.Columns[index].Category = CAT_FLOAT64
	} else if strings.Contains(columnType, "TEXT") {
		ta.Columns[index].Category = CAT_STR
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
