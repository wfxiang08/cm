// Copyright 2014, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package planbuilder

import (
	"github.com/juju/errors"

	"github.com/ngaut/arena"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/sqlparser"
	"github.com/wandoulabs/cm/vt/schema"
)

var (
	TooComplex = errors.New("Complex")
	execLimit  = &sqlparser.Limit{Rowcount: sqlparser.ValArg(":#maxLimit")}
)

// ExecPlan is built for selects and DMLs.
// PK Values values within ExecPlan can be:
// sqltypes.Value: sourced form the query, or
// string: bind variable name starting with ':', or
// nil if no value was specified
type ExecPlan struct {
	PlanId    PlanType
	Reason    ReasonType
	TableName string

	// FieldQuery is used to fetch field info
	FieldQuery *sqlparser.ParsedQuery

	// FullQuery will be set for all plans.
	FullQuery *sqlparser.ParsedQuery

	// For PK plans, only OuterQuery is set.
	// For SUBQUERY plans, Subquery is also set.
	// IndexUsed is set only for PLAN_SELECT_SUBQUERY
	OuterQuery *sqlparser.ParsedQuery
	Subquery   *sqlparser.ParsedQuery
	IndexUsed  string

	// For selects, columns to be returned
	// For PLAN_INSERT_SUBQUERY, columns to be inserted
	ColumnNumbers []int

	// PLAN_PK_IN, PLAN_DML_PK: where clause values
	// PLAN_INSERT_PK: values clause
	PKValues []interface{}

	// PK_IN. Limit clause value.
	Limit interface{}

	// For update: set clause if pk is changing
	SecondaryPKValues []interface{}

	// For PLAN_INSERT_SUBQUERY: pk columns in the subquery result
	SubqueryPKColumns []int

	// PLAN_SET
	SetKey   string
	SetValue interface{}
}

func (node *ExecPlan) setTableInfo(tableName string, getTable TableGetter) (*schema.Table, error) {
	tableInfo, ok := getTable(tableName)
	if !ok {
		return nil, errors.Errorf("table %s not found in schema", tableName)
	}
	node.TableName = tableInfo.Name
	return tableInfo, nil
}

type TableGetter func(tableName string) (*schema.Table, bool)

// func GetSqlExecPlan(sql string, getTable TableGetter, alloc arena.ArenaAllocator) (plan *ExecPlan, err error) {
// statement, err := sqlparser.Parse(sql, alloc)
// if err != nil {
// return nil, err
// }
// plan, err = analyzeSQL(statement, getTable, alloc)
// if err != nil {
// return nil, err
// }
// if plan.PlanId == PLAN_PASS_DML {
// log.Warningf("PASS_DML: %s", sql)
// }
// return plan, nil
// }

func GetStmtExecPlan(stmt sqlparser.Statement, getTable TableGetter, alloc arena.ArenaAllocator) (plan *ExecPlan, err error) {
	plan, err = analyzeSQL(stmt, getTable, alloc)
	if err != nil {
		return nil, err
	}

	if plan.PlanId == PLAN_PASS_DML {
		log.Warningf("PASS_DML: %s", sqlparser.String(stmt, alloc))
	}

	return plan, nil
}

func analyzeSQL(statement sqlparser.Statement, getTable TableGetter, alloc arena.ArenaAllocator) (plan *ExecPlan, err error) {
	switch stmt := statement.(type) {
	case *sqlparser.Union:
		return &ExecPlan{
			PlanId:     PLAN_PASS_SELECT,
			FieldQuery: GenerateFieldQuery(stmt, alloc),
			FullQuery:  GenerateFullQuery(stmt, alloc),
			Reason:     REASON_SELECT,
		}, nil
	case *sqlparser.Select:
		return analyzeSelect(stmt, getTable, alloc)
	case *sqlparser.Insert:
		return analyzeInsert(stmt, getTable, alloc)
	case *sqlparser.Replace:
		return analyzeReplace(stmt, getTable, alloc)
	case *sqlparser.Update:
		return analyzeUpdate(stmt, getTable, alloc)
	case *sqlparser.Delete:
		return analyzeDelete(stmt, getTable, alloc)
	case *sqlparser.Set:
		return analyzeSet(stmt, alloc), nil
	case *sqlparser.DDL:
		return analyzeDDL(stmt, getTable), nil
	case *sqlparser.Other:
		return &ExecPlan{PlanId: PLAN_OTHER}, nil
	}
	return nil, errors.New("invalid SQL")
}
