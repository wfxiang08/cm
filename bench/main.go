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
	N              = flag.Int("n", 10000, "test row count")
	C              = flag.Int("c", 0, "concurrent worker")
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
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s?useServerPrepStmts=false&charset=utf8mb4", *mysqlProxyHost, *mysqlProxyPort, *dbName)
	return NewDb(dsn)
}

func NewProxyDbWithCharset(charsetName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s?useServerPrepStmts=false&charset=%s", *mysqlProxyHost, *mysqlProxyPort, *dbName, charsetName)
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

	dropBenchTbls(mysqlDb)
	createBenchTbls(mysqlDb)

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
}
