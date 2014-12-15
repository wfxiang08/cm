package mysql

import (
	"fmt"
	"strconv"

	"github.com/wandoulabs/cm/hack"
)

type RowData []byte
type Value interface{}
type RowValue []Value

func Raw(t byte, val Value, isUnsigned bool) []byte {
	var ret []byte
	switch t {
	case MYSQL_TYPE_TINY, MYSQL_TYPE_SHORT, MYSQL_TYPE_INT24,
		MYSQL_TYPE_LONGLONG, MYSQL_TYPE_YEAR:
		if isUnsigned {
			ret = []byte(strconv.FormatUint(val.(uint64), 10))
		} else {
			ret = []byte(strconv.FormatInt(val.(int64), 10))
		}

	case MYSQL_TYPE_FLOAT, MYSQL_TYPE_DOUBLE:
		ret = []byte(strconv.FormatFloat(val.(float64), 'f', 7, 10))
	default:
		ret = val.([]byte)
	}
	return ret
}

func (p RowData) Parse(f []*Field, binary bool) (RowValue, error) {
	return p.ParseText(f)
}

func (p RowData) ParseText(f []*Field) (RowValue, error) {
	data := make([]Value, len(f))

	var err error
	var v []byte
	var isNull, isUnsigned bool
	var pos int = 0
	var n int = 0

	for i := range f {
		v, isNull, n, err = LengthEnodedString(p[pos:])
		if err != nil {
			return nil, err
		}

		pos += n

		if isNull {
			data[i] = nil
		} else {
			isUnsigned = (f[i].Flag&UNSIGNED_FLAG > 0)

			switch f[i].Type {
			case MYSQL_TYPE_TINY, MYSQL_TYPE_SHORT, MYSQL_TYPE_INT24,
				MYSQL_TYPE_LONGLONG, MYSQL_TYPE_YEAR:
				if isUnsigned {
					data[i], err = strconv.ParseUint(string(v), 10, 64)
				} else {
					data[i], err = strconv.ParseInt(string(v), 10, 64)
				}
			case MYSQL_TYPE_FLOAT, MYSQL_TYPE_DOUBLE:
				data[i], err = strconv.ParseFloat(string(v), 64)
			default:
				data[i] = v
			}

			if err != nil {
				return nil, err
			}
		}
	}

	return data, nil
}

type Resultset struct {
	Fields     []*Field
	FieldNames map[string]int
	Values     []RowValue

	RowDatas []RowData
}

func (r *Resultset) RowNumber() int {
	return len(r.Values)
}

func (r *Resultset) ColumnNumber() int {
	return len(r.Fields)
}

func (r *Resultset) GetValue(row, column int) (interface{}, error) {
	if row >= len(r.Values) || row < 0 {
		return nil, fmt.Errorf("invalid row index %d", row)
	}

	if column >= len(r.Fields) || column < 0 {
		return nil, fmt.Errorf("invalid column index %d", column)
	}

	return r.Values[row][column], nil
}

func (r *Resultset) NameIndex(name string) (int, error) {
	if column, ok := r.FieldNames[name]; ok {
		return column, nil
	} else {
		return 0, fmt.Errorf("invalid field name %s", name)
	}
}

func (r *Resultset) GetValueByName(row int, name string) (interface{}, error) {
	if column, err := r.NameIndex(name); err != nil {
		return nil, err
	} else {
		return r.GetValue(row, column)
	}
}

func (r *Resultset) IsNull(row, column int) (bool, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return false, err
	}

	return d == nil, nil
}

func (r *Resultset) IsNullByName(row int, name string) (bool, error) {
	if column, err := r.NameIndex(name); err != nil {
		return false, err
	} else {
		return r.IsNull(row, column)
	}
}

func (r *Resultset) GetUint(row, column int) (uint64, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return 0, err
	}

	switch v := d.(type) {
	case uint64:
		return v, nil
	case int64:
		return uint64(v), nil
	case float64:
		return uint64(v), nil
	case string:
		return strconv.ParseUint(v, 10, 64)
	case []byte:
		return strconv.ParseUint(string(v), 10, 64)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("data type is %T", v)
	}
}

func (r *Resultset) GetUintByName(row int, name string) (uint64, error) {
	if column, err := r.NameIndex(name); err != nil {
		return 0, err
	} else {
		return r.GetUint(row, column)
	}
}

func (r *Resultset) GetInt(row, column int) (int64, error) {
	v, err := r.GetUint(row, column)
	if err != nil {
		return 0, err
	}

	return int64(v), nil
}

func (r *Resultset) GetIntByName(row int, name string) (int64, error) {
	v, err := r.GetUintByName(row, name)
	if err != nil {
		return 0, err
	}

	return int64(v), nil
}

func (r *Resultset) GetFloat(row, column int) (float64, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return 0, err
	}

	switch v := d.(type) {
	case float64:
		return v, nil
	case uint64:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	case []byte:
		return strconv.ParseFloat(string(v), 64)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("data type is %T", v)
	}
}

func (r *Resultset) GetFloatByName(row int, name string) (float64, error) {
	if column, err := r.NameIndex(name); err != nil {
		return 0, err
	} else {
		return r.GetFloat(row, column)
	}
}

func (r *Resultset) GetString(row, column int) (string, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return "", err
	}

	switch v := d.(type) {
	case string:
		return v, nil
	case []byte:
		return hack.String(v), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("data type is %T", v)
	}
}

func (r *Resultset) GetStringByName(row int, name string) (string, error) {
	if column, err := r.NameIndex(name); err != nil {
		return "", err
	} else {
		return r.GetString(row, column)
	}
}
