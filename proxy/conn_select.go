package proxy

import (
	"bytes"
	"strings"

	"github.com/juju/errors"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/hack"
	. "github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/sqlparser"
)

func (c *Conn) handleSimpleSelect(sql string, stmt *sqlparser.SimpleSelect) error {
	if len(stmt.SelectExprs) != 1 {
		return errors.Errorf("support select one informaction function, %s", sql)
	}

	expr, ok := stmt.SelectExprs[0].(*sqlparser.NonStarExpr)
	if !ok {
		return errors.Errorf("support select informaction function, %s", sql)
	}

	var funcExpr *sqlparser.FuncExpr
	var specialColumn *sqlparser.ColName

	switch v := expr.Expr.(type) {
	case *sqlparser.FuncExpr:
		funcExpr = v
	case *sqlparser.ColName:
		specialColumn = v
		log.Debug(specialColumn)
	case sqlparser.NumVal:
		return errors.Trace(c.handleShow(stmt, sql, nil))
	case *sqlparser.BinaryExpr:
		return errors.Trace(c.handleShow(stmt, sql, nil))
	default:
		return errors.Errorf("support select informaction function, %s, %T", sql, v)
	}

	var r *Resultset
	var err error

	if funcExpr != nil {
		switch strings.ToLower(string(funcExpr.Name)) {
		case "last_insert_id":
			r, err = c.buildSimpleSelectResult(c.lastInsertId, funcExpr.Name, expr.As)
		case "row_count":
			r, err = c.buildSimpleSelectResult(c.affectedRows, funcExpr.Name, expr.As)
		case "version":
			r, err = c.buildSimpleSelectResult(ServerVersion, funcExpr.Name, expr.As)
		case "connection_id":
			r, err = c.buildSimpleSelectResult(c.connectionId, funcExpr.Name, expr.As)
		case "database":
			if c.schema != nil {
				r, err = c.buildSimpleSelectResult(c.schema.db, funcExpr.Name, expr.As)
			} else {
				r, err = c.buildSimpleSelectResult("NULL", funcExpr.Name, expr.As)
			}
		default:
			return errors.Errorf("function %s not support", funcExpr.Name)
		}

		if err != nil {
			return errors.Trace(err)
		}
	} else {
		return errors.Trace(c.handleShow(stmt, sql, nil))
	}

	return errors.Trace(c.writeResultset(c.status, r))
}

func (c *Conn) buildSimpleSelectResult(value interface{}, name []byte, asName []byte) (*Resultset, error) {
	field := &Field{Name: name, OrgName: name}
	if asName != nil {
		field.Name = asName
	}

	formatField(field, value)

	r := &Resultset{Fields: []*Field{field}}
	row := Raw(byte(field.Type), value, false)
	r.RowDatas = append(r.RowDatas, PutLengthEncodedString(row))

	return r, nil
}

func (c *Conn) handleFieldList(data []byte) error {
	index := bytes.IndexByte(data, 0x00)
	table := hack.String(data[0:index])
	wildcard := hack.String(data[index+1:])

	if c.schema == nil {
		return errors.Trace(NewDefaultError(ER_NO_DB_ERROR))
	}

	nodeName := c.schema.rule.GetRule(table).Node
	//todo: pass through
	if len(nodeName) == 0 {
		return errors.Errorf("no rule for table %s, %+v, please check config file", table, c.schema)
	}

	n := c.server.getNode(nodeName)
	if n == nil {
		return errors.Errorf("node %s not found, %+v", nodeName, c.schema)
	}

	co, err := n.getMasterConn()
	if err != nil {
		return errors.Trace(err)
	}
	defer co.Close()

	if err = co.UseDB(c.schema.db); err != nil {
		return errors.Trace(err)
	}

	if fs, err := co.FieldList(table, wildcard); err != nil {
		return errors.Trace(err)
	} else {
		return errors.Trace(c.writeFieldList(c.status, fs))
	}
}

func (c *Conn) writeFieldList(status uint16, fs []*Field) error {
	c.affectedRows = int64(-1)

	data := make([]byte, 4, 1024)

	for _, v := range fs {
		data = data[0:4]
		data = append(data, v.Dump()...)
		if err := c.writePacket(data); err != nil {
			return errors.Trace(err)
		}
	}

	err := c.writeEOF(status)
	return errors.Trace(err)
}
