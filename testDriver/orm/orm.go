package main

import (
	"database/sql"
	"flag"
	"math/rand"
	"runtime"
	"sync"

	_ "github.com/c4pt0r/mysql"

	"github.com/coopernurse/gorp"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewDbMap() *gorp.DbMap {
	dsn := "root:@tcp(127.0.0.1:4000)/test"
	dbType := "mysql"
	db, err := sql.Open(dbType, dsn)

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	if err != nil {
		panic(err.Error())
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(TestData{}, "autoincr_test").SetKeys(true, "Id")

	if err != nil {
		panic(err.Error())
	}
	return dbmap
}

type TestData struct {
	Id       int    `db:"id"`
	Data     string `db:"data"`
	Datetime []byte `db:"datetime"`
}

var n = flag.Int("n", 20000, "n")
var testType = flag.String("t", "read", "read|write")
var numWorkers = 50

func WriteTest() {
	wg := sync.WaitGroup{}
	var chans []chan *TestData
	for i := 0; i < numWorkers; i++ {
		c := make(chan *TestData, 100)
		chans = append(chans, c)
		go func(chan *TestData) {
			m := NewDbMap()
			for t := range c {
				err := m.Insert(t)
				if err != nil {
					panic(err)
				}
				wg.Done()
			}
		}(c)
	}
	for i := 0; i < *n; i++ {
		wg.Add(1)
		chans[i%numWorkers] <- &TestData{
			Data: randSeq(1024),
		}
	}
	wg.Wait()
}

func ReadTest() {
	wg := sync.WaitGroup{}
	c := make(chan int)
	for i := 0; i < numWorkers; i++ {
		go func(c chan int) {
			m := NewDbMap()
			t := TestData{}
			for _ = range c {
				err := m.SelectOne(&t, "select * from autoincr_test where id = ?", 1)
				if err != nil {
					panic(err)
				}
				wg.Done()
			}
		}(c)
	}

	for i := 1; i < *n; i++ {
		wg.Add(1)
		c <- i
	}

	wg.Wait()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	if *testType == "read" {
		ReadTest()
	} else {
		WriteTest()
	}
}
