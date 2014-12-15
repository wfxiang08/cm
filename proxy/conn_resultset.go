package proxy

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/wandoulabs/cm/hack"
	. "github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/vt/schema"
)

func formatField(field *Field, value interface{}) error {
	switch value.(type) {
	case int8, int16, int32, int64, int:
		field.Charset = 63
		field.Type = MYSQL_TYPE_LONGLONG
		field.Flag = BINARY_FLAG | NOT_NULL_FLAG
	case uint8, uint16, uint32, uint64, uint:
		field.Charset = 63
		field.Type = MYSQL_TYPE_LONGLONG
		field.Flag = BINARY_FLAG | NOT_NULL_FLAG | UNSIGNED_FLAG
	case string, []byte:
		field.Charset = 33
		field.Type = MYSQL_TYPE_VARCHAR
	case nil:
		return nil
	default:
		return fmt.Errorf("unsupport type %T for resultset", value)
	}
	return nil
}

func (c *Conn) buildResultset(nameTypes []schema.TableColumn, values []RowValue) (*Resultset, error) {
	r := &Resultset{Fields: make([]*Field, len(nameTypes))}

	var b []byte
	var err error

	for i, vs := range values {
		if len(vs) != len(r.Fields) {
			return nil, fmt.Errorf("row %d has %d column not equal %d", i, len(vs), len(r.Fields))
		}

		var row []byte
		for j, value := range vs {
			field := &Field{}
			if i == 0 {
				r.Fields[j] = field
				field.Name = hack.Slice(nameTypes[j].Name)
				if err = formatField(field, value); err != nil {
					return nil, errors.Trace(err)
				}
				field.Type = nameTypes[j].Category
			}

			if value == nil {
				row = append(row, "\xfb"...)
			} else {
				b = Raw(byte(field.Type), value, false)
				row = append(row, PutLengthEncodedString(b)...)
			}
		}

		r.RowDatas = append(r.RowDatas, row)
	}

	return r, nil
}

func (c *Conn) writeResultset(status uint16, r *Resultset) error {
	c.affectedRows = int64(-1)
	columnLen := PutLengthEncodedInt(uint64(len(r.Fields)))
	data := make([]byte, 4, 1024)
	data = append(data, columnLen...)
	if err := c.writePacket(data); err != nil {
		return errors.Trace(err)
	}

	for _, v := range r.Fields {
		data = data[0:4]
		data = append(data, v.Dump()...)
		if err := c.writePacket(data); err != nil {
			return errors.Trace(err)
		}
	}

	if err := c.writeEOF(status); err != nil {
		return errors.Trace(err)
	}

	for _, v := range r.RowDatas {
		data = data[0:4]
		data = append(data, v...)
		if err := c.writePacket(data); err != nil {
			return errors.Trace(err)
		}
	}

	err := c.writeEOF(status)
	return errors.Trace(err)
}
