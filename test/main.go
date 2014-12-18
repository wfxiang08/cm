package main

import (
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
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
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s?useServerPrepStmts=true", *mysqlHost, *mysqlPort, *dbName)
	return NewDb(dsn)
}

func NewProxyDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s", *mysqlProxyHost, *mysqlProxyPort, *dbName)
	return NewDb(dsn)
}

func insertAutoIncrIdData(db *sql.DB, data string) (int64, error) {
	return -1, fmt.Errorf("not implement")
}

/* insert test */
func testInsertData(db *sql.DB) error {
	for i := 0; i < *N; i++ {
		_, err := insertAutoIncrIdData(db, randSeq(100))
		if err != nil {
			return err
		}
	}
	return nil
}

func testConcurrentInsert(db *sql.DB) error {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	e := make(chan error)
	for i := 0; i < *C; i++ {
		go func() {
			for _ = range c {
				_, err := insertAutoIncrIdData(db, randSeq(100))
				if err != nil {
					e <- err
				}
				wg.Done()
			}
		}()
	}

	wg.Add(*N)
	for i := 0; i < *N; i++ {
		c <- struct{}{}
	}
	wg.Wait()
	return nil
}

/* read test */

func testRead(db *sql.DB) error {
	for i := 0; i < *N; i++ {
		x := rand.Intn(2000) + 1
		r, err := db.Query("select * from autoincr_test where id = ?", x)
		if err != nil {
			return err
		}
		r.Close()
	}
	return nil
}

func testConcurrentRead(db *sql.DB) error {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	e := make(chan error)
	for i := 0; i < *C; i++ {
		go func() {
			for _ = range c {
				x := rand.Intn(2000) + 1
				r, err := db.Query("select * from autoincr_test where id = ?", x)
				if err != nil {
					e <- err
				}
				r.Close()
				wg.Done()
			}
		}()
	}

	wg.Add(*N)
	for i := 0; i < *N; i++ {
		c <- struct{}{}
	}
	wg.Wait()
	return nil
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

	case "write":
		if *C == 0 {
			fmt.Println(run("write test", db, testInsertData))
		} else {
			fmt.Println(run("concurrent write test", db, testConcurrentInsert))
		}
	case "read":
		if *C == 0 {
			fmt.Println(run("read test", db, testRead))
		} else {
			fmt.Println(run("concurrent read test", db, testConcurrentRead))
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

	default:
		fmt.Printf("no such type")
	}
}
