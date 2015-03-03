package test

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
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

var proxyDB *sql.DB
var mysqlDB *sql.DB

func newDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	return db, nil
}

func newMysqlDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s", mysqlHost, mysqlPort, dbName)
	return newDb(dsn)
}

func newProxyDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s?useServerPrepStmts=false", mysqlProxyHost, mysqlProxyPort, dbName)
	return newDb(dsn)
}

func newProxyDbWithCharset(charsetName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s?useServerPrepStmts=false&charset=%s", mysqlProxyHost, mysqlProxyPort, dbName, charsetName)
	return newDb(dsn)
}

func init() {
	var err error
	proxyDB, err = newProxyDb()
	if err != nil {
		panic(err)
	}
	mysqlDB, err = newMysqlDb()
	if err != nil {
		panic(err)
	}
}

func dropTables() {
	doEachTbl(func(tblName string) {
		mustExec(mysqlDB, "drop table "+tblName)
	})
}

func doEachTbl(fn func(tblName string)) {
	db := mysqlDB
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			log.Fatal(err)
		}
		fn(table)
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

type SqlExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func mustExec(execer SqlExecer, sql string, args ...interface{}) sql.Result {
	res, err := execer.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
	return res
}

func mustQueryDataWithMultiId(db *sql.DB, tblName string, id1 interface{}, id2 interface{}, data interface{}) {
	rows := mustQuery(proxyDB, "select data from "+tblName+" where id1=? and id2=?", id1, id2)
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

func queryData(db *sql.DB, tblName string, id interface{}, data interface{}) error {
	rows := mustQuery(proxyDB, "select data from "+tblName+" where id=?", id)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(data)
		if err != nil {
			return err
		}
	}

	return nil
}

func mustQueryData(db *sql.DB, tblName string, id interface{}, data interface{}) {
	rows := mustQuery(proxyDB, "select data from "+tblName+" where id=?", id)
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
	db := mysqlDB
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
	case uint64:
		return v == b.(uint64)
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
