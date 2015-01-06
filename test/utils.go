package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"reflect"

	_ "github.com/c4pt0r/mysql"
)

var (
	mysqlProxyHost = "127.0.0.1"
	mysqlProxyPort = 4000
	mysqlHost      = "127.0.0.1"
	mysqlPort      = 3306
	dbName         = "test"
)

var ProxyDB *sql.DB
var MysqlDB *sql.DB

func NewDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	return db, nil
}

func NewMysqlDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s", mysqlHost, mysqlPort, dbName)
	return NewDb(dsn)
}

func NewProxyDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s?useServerPrepStmts=false", mysqlProxyHost, mysqlProxyPort, dbName)
	return NewDb(dsn)
}

func NewProxyDbWithCharset(charsetName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s?useServerPrepStmts=false&charset=%s", mysqlProxyHost, mysqlProxyPort, dbName, charsetName)
	return NewDb(dsn)
}

func init() {
	var err error
	ProxyDB, err = NewProxyDb()
	if err != nil {
		panic(err)
	}
	MysqlDB, err = NewMysqlDb()
	if err != nil {
		panic(err)
	}
}

func reloadConfig() {
	http.Get("http://127.0.0.1:8888/api/reload")
}

func mustQuery(db *sql.DB, sql string, args ...interface{}) *sql.Rows {
	res, err := db.Query(sql, args...)
	if err != nil {
		panic(err)
	}
	return res
}

func mustExec(db *sql.DB, sql string, args ...interface{}) sql.Result {
	res, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
	return res
}

func mustQueryData(db *sql.DB, tblName string, id interface{}, data interface{}) {
	rows := mustQuery(ProxyDB, "select data from "+tblName+" where id=?", id)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(data)
		if err != nil {
			panic(err)
		}
	} else {
		panic("no such record")
	}
}

func isTblExists(tblName string) bool {
	db := MysqlDB
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			panic(err)
		}
		if table == tblName {
			return true
		}
	}
	return false
}

func equal(a interface{}, b interface{}) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}
	switch v := a.(type) {
	case int:
		return v == b.(int)
	case int64:
		return v == b.(int64)
	case string:
		return v == b.(string)
	case []byte:
		return bytes.Equal(v, b.([]byte))
	case float64:
		return math.Abs(v-b.(float64)) < 1e-7
	default:
		panic("not supported compare")
	}
}
