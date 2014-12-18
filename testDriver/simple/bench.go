package main

import (
	"database/sql"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"

	_ "github.com/c4pt0r/mysql"
)

var (
	mysqlHost = flag.String("h", "127.0.0.1", "mysql host")
	mysqlPort = flag.Int("p", 4000, "mysql port")
	dbName    = flag.String("db", "test", "db name")

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

func createTestTbls(db *sql.DB) {
	if b, err := isTblExists(db, "autoincr_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE autoincr_test(id INT NOT NULL AUTO_INCREMENT, data VARCHAR(1024), datetime DATETIME, primary KEY(id))`)
	}

	if b, err := isTblExists(db, "int_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE int_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data INT)`)
	}

	if b, err := isTblExists(db, "double_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE double_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data DOUBLE)`)
	}

	if b, err := isTblExists(db, "varchar_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE varchar_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data VARCHAR(1024))`)
	}

	if b, err := isTblExists(db, "text_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE text_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data TEXT)`)
	}

	if b, err := isTblExists(db, "blob_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE blob_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data BLOB)`)
	}

	if b, err := isTblExists(db, "datetime_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE datetime_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data DATETIME)`)
	}

	if b, err := isTblExists(db, "date_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE date_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data DATE)`)
	}
}

func dropTestTbls(db *sql.DB) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var tbls []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			panic(err)
		}
		tbls = append(tbls, table)
	}

	for _, t := range tbls {
		mustExec(db, "DROP TABLE "+t)
	}
}

func insertAutoIncrIdData(db *sql.DB, data string) (int64, error) {
	res := mustExec(db, "INSERT INTO autoincr_test(data) VALUES(?)", data)
	return res.LastInsertId()
}

func insertDataAndQueryBack(db *sql.DB, tblName string, val interface{}, ret interface{}, cachedRet interface{}) error {
	r := mustExec(db, "INSERT INTO "+tblName+"(data) VALUES(?)", val)
	id, err := r.LastInsertId()
	if err != nil {
		return err
	}

	res, err := db.Query("SELECT data FROM "+tblName+" WHERE id = ?", id)
	if err != nil {
		return err
	}
	defer res.Close()

	for res.Next() {
		res.Scan(ret)
		if err != nil {
			return err
		}
	}

	cachedRes, err := db.Query("SELECT data FROM "+tblName+" WHERE id = ?", id)
	if err != nil {
		return err
	}
	defer cachedRes.Close()

	for cachedRes.Next() {
		cachedRes.Scan(cachedRet)
		if err != nil {
			return err
		}
	}

	return nil

}

func intTest(db *sql.DB) error {
	x := rand.Intn(1024) + 1
	var ret, cret int
	err := insertDataAndQueryBack(db, "int_test", x, &ret, &cret)
	if err != nil {
		return err
	}
	if ret != x {
		return fmt.Errorf("int test failed %d != %d", ret, x)
	}
	return nil
}

func doubleTest(db *sql.DB) error {
	x := rand.Float64()
	var ret, cret float64

	err := insertDataAndQueryBack(db, "double_test", x, &ret, &cret)
	if err != nil {
		return err
	}

	if math.Abs(ret-x) > 1e-7 {
		return fmt.Errorf("double test failed %v != %v", ret, x)
	}
	return nil
}

func varcharTest(db *sql.DB) error {
	s := randSeq(1024)
	var ret, cret string
	err := insertDataAndQueryBack(db, "varchar_test", s, &ret, &cret)
	if err != nil {
		return err
	}

	if ret != s {
		return fmt.Errorf("var char test failed %v != %v", ret, s)
	}
	return nil
}

func textTest(db *sql.DB) error {
	s := randSeq(4096)
	var ret, cret string
	err := insertDataAndQueryBack(db, "text_test", s, &ret, &cret)
	if err != nil {
		return err
	}

	if ret != s {
		return fmt.Errorf("text test failed %v != %v", ret, s)
	}
	return nil
}

func blobTest(db *sql.DB) error {
	blob := []byte{
		0xff, 0x00, 0x01, 0x02, 0x03,
	}
	var ret, cret []byte
	err := insertDataAndQueryBack(db, "blob_test", blob, &ret, &cret)
	if err != nil {
		return err
	}

	s1 := fmt.Sprintf("%x", ret)
	s2 := fmt.Sprintf("%x", blob)
	if s1 != s2 {
		return fmt.Errorf("blob test failed %v != %v", s1, s2)
	}
	return nil
}

func datetimeTest(db *sql.DB) error {
	d := time.Now()
	ds := d.Format("2006-01-02 03:04:05")
	var ret, cret []byte
	err := insertDataAndQueryBack(db, "datetime_test", d, &ret, &cret)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	if string(ret) != ds {
		return fmt.Errorf("datetime test failed %v != %v", string(ret), ds)
	}

	return nil
}
func dateTest(db *sql.DB) error {
	d := time.Now()
	ds := d.Format("2006-01-02")
	var ret, cret []byte
	err := insertDataAndQueryBack(db, "date_test", d, &ret, &cret)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	if string(ret) != ds {
		return fmt.Errorf("date test failed %v != %v", string(ret), ds)
	}

	return nil
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
	db, err := NewDb()
	rand.Seed(time.Now().UnixNano())

	if err != nil {
		panic(err)
	}

	go func() {
		http.ListenAndServe(":8889", nil)
	}()

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
		if *C == 0 {
			fmt.Println(run(db, testRead))
		} else {
			fmt.Println(run(db, testConcurrentRead))
		}
	case "type-test":
		fmt.Println(run(db, intTest))
		fmt.Println(run(db, doubleTest))
		fmt.Println(run(db, varcharTest))
		fmt.Println(run(db, textTest))
		fmt.Println(run(db, blobTest))
		fmt.Println(run(db, datetimeTest))
		fmt.Println(run(db, dateTest))

	default:
		fmt.Printf("no such type")
	}
}
