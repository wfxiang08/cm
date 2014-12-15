package main

import (
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	_ "github.com/c4pt0r/mysql"
)

var (
	mysqlHost = flag.String("h", "127.0.0.1", "mysql host")
	mysqlPort = flag.Int("p", 4000, "mysql port")
	dbName    = flag.String("db", "benchmark", "db name")

	testType = flag.String("t", "read", "test type: read | write | init")
	N        = flag.Int("n", 10000, "test row count")
	C        = flag.Int("c", 0, "concurrent worker")
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type Result struct {
	Queries  int
	Duration time.Duration
	Allocs   uint64
	Bytes    uint64
}

func (res Result) String() string {
	return fmt.Sprintln(
		" "+res.Duration.String(), "\t   ",
		int(res.QueriesPerSecond()+0.5), "queries/sec\t   ",
		res.AllocsPerQuery(), "allocs/query\t   ",
		res.BytesPerQuery(), "B/query",
	)
}

func (res Result) QueriesPerSecond() float64 {
	return float64(res.Queries) / res.Duration.Seconds()
}

func (res Result) AllocsPerQuery() int {
	return int(res.Allocs) / res.Queries
}

func (res Result) BytesPerQuery() int {
	return int(res.Bytes) / res.Queries
}

var memStats runtime.MemStats

func run(db *sql.DB, testFunc func(*sql.DB) error) Result {
	runtime.GC()

	runtime.ReadMemStats(&memStats)
	var (
		startMallocs    = memStats.Mallocs
		startTotalAlloc = memStats.TotalAlloc
		startTime       = time.Now()
	)

	err := testFunc(db)
	if err != nil {
		panic(err)
	}

	endTime := time.Now()
	runtime.ReadMemStats(&memStats)

	return Result{
		Queries:  *N,
		Duration: endTime.Sub(startTime),
		Allocs:   memStats.Mallocs - startMallocs,
		Bytes:    memStats.TotalAlloc - startTotalAlloc,
	}
}

func NewDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("root:@tcp(%s:%d)/%s", *mysqlHost, *mysqlPort, *dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	return db, nil
}

func mustExec(db *sql.DB, sql string, args ...interface{}) sql.Result {
	res, err := db.Exec(sql, args...)
	if err != nil {
		panic(err)
	}
	return res
}

func isTblExists(db *sql.DB, tblName string) (bool, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return false, err
	}
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			return false, err
		}
		if table == tblName {
			return true, nil
		}
	}
	return false, nil
}

func createTestTbls(db *sql.DB) {
	if b, err := isTblExists(db, "autoincr_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE autoincr_test(id INT NOT NULL AUTO_INCREMENT, data VARCHAR(1024), datetime DATETIME, primary KEY(id))`)
	}
}

func dropTestTbls(db *sql.DB) {
	if b, err := isTblExists(db, "autoincr_test"); b && err == nil {
		mustExec(db, `DROP TABLE autoincr_test`)
	}
}

func insertAutoIncrIdData(db *sql.DB, data string) (int64, error) {
	res := mustExec(db, "INSERT INTO autoincr_test(data) VALUES(?)", data)
	return res.LastInsertId()
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
		_, err := db.Query("select * from autoincr_test where id = ?", x)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	db, err := NewDb()

	if err != nil {
		panic(err)
	}

	switch *testType {
	case "init":
		dropTestTbls(db)
		createTestTbls(db)
	case "write":
		if *C == 0 {
			fmt.Println(run(db, testInsertData))
		} else {
			fmt.Println(run(db, testConcurrentInsert))
		}
	case "read":
		fmt.Println(run(db, testRead))
	default:
		fmt.Printf("no such type")
	}
}
