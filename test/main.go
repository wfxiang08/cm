package main

import (
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"

	_ "github.com/c4pt0r/mysql"
)

var (
	mysqlProxyHost = flag.String("h", "127.0.0.1", "mysql proxy host")
	mysqlProxyPort = flag.Int("p", 4000, "mysql proxy port")
	mysqlHost      = flag.String("H", "127.0.0.1", "mysql host")
	mysqlPort      = flag.Int("P", 3306, "mysql port")
	dbName         = flag.String("db", "test", "db name")

	testType = flag.String("t", "read", "test type: read | write | init")
	N        = flag.Int("n", 10000, "test row count")
	C        = flag.Int("c", 0, "concurrent worker")
)

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
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s", *mysqlHost, *mysqlPort, *dbName)
	return NewDb(dsn)
}

func NewProxyDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s?useServerPrepStmts=false", *mysqlProxyHost, *mysqlProxyPort, *dbName)
	return NewDb(dsn)
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	db, err := NewProxyDb()
	if err != nil {
		panic(err)
	}

	mysqlDb, err := NewMysqlDb()
	if err != nil {
		panic(err)
	}

	go func() {
		http.ListenAndServe(":8889", nil)
	}()

	switch *testType {
	case "init":
		dropTypeTestTbls(mysqlDb)
		createTypeTestTbls(mysqlDb)
		dropBenchTbls(mysqlDb)
		createBenchTbls(mysqlDb)
	case "bench":
		createBenchTbls(mysqlDb)
		defer dropBenchTbls(mysqlDb)

		fmt.Println(run("insert bench data", db, insertBenchData).String())
		if *C > 0 {
			fmt.Println(run("insert bench data async", db, concurrentInsertBenchData).String())
		}
		fmt.Println(run("read bench data", db, readBenchData).String())
		if *C > 0 {
			fmt.Println(run("read bench data async", db, concurrentReadBenchData).String())
		}
		fmt.Println(run("read bench data on cache", db, readBenchData).String())
		if *C > 0 {
			fmt.Println(run("read bench data async on cache", db, concurrentReadBenchData).String())
		}

	case "type-test":
		createTypeTestTbls(mysqlDb)
		defer dropTypeTestTbls(mysqlDb)
		fmt.Printf(run("int type test", db, intTest).String())
		fmt.Printf(run("double type test", db, doubleTest).String())
		fmt.Printf(run("varchar type test", db, varcharTest).String())
		fmt.Printf(run("text type test", db, textTest).String())
		fmt.Printf(run("blob type test", db, blobTest).String())
		fmt.Printf(run("datetime type test", db, datetimeTest).String())
		fmt.Printf(run("date type test", db, dateTest).String())
		fmt.Printf(run("multi pkey test", db, multiPKeyTest).String())

	default:
		fmt.Printf("no such type")
	}
}
