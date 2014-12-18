package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	_ "github.com/c4pt0r/mysql"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// random string
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// simple test run suit
type Result struct {
	Desc     string
	Err      error
	Queries  int
	Duration time.Duration
	Allocs   uint64
	Bytes    uint64
}

func (res Result) String() string {
	runResult := "OK"
	if res.Err != nil {
		runResult = "ERR: " + res.Err.Error()
	}
	return fmt.Sprintln(
		res.Desc+"\t"+runResult+":\n"+res.Duration.String(), "\t   ",
		int(res.QueriesPerSecond()+0.5), "queries/sec\t   ",
		res.AllocsPerQuery(), "allocs/query\t   ",
		res.BytesPerQuery(), "B/query\n",
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

func run(testName string, db *sql.DB, testFunc func(*sql.DB) error) Result {
	runtime.GC()

	runtime.ReadMemStats(&memStats)
	var (
		startMallocs    = memStats.Mallocs
		startTotalAlloc = memStats.TotalAlloc
		startTime       = time.Now()
	)

	err := testFunc(db)

	endTime := time.Now()
	runtime.ReadMemStats(&memStats)

	return Result{
		Desc:     testName,
		Err:      err,
		Queries:  *N,
		Duration: endTime.Sub(startTime),
		Allocs:   memStats.Mallocs - startMallocs,
		Bytes:    memStats.TotalAlloc - startTotalAlloc,
	}
}

// mysql utils
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
	defer rows.Close()
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
