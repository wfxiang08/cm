package main

import (
	"database/sql"
	"log"
	"math/rand"
	"sync"

	_ "github.com/c4pt0r/mysql"
)

func createBenchTbls(db *sql.DB) {
	if b, err := isTblExists(db, "tbl_bench"); !b && err == nil {
		mustExec(db, `CREATE TABLE tbl_bench(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data VARCHAR(1024))`)
	} else if err != nil {
		log.Fatal(err)
	}
}

func dropBenchTbls(db *sql.DB) {
	tbls := []string{
		"tbl_bench",
	}

	for _, t := range tbls {
		if b, _ := isTblExists(db, t); b {
			mustExec(db, "DROP TABLE "+t)
		}
	}
}

func insertAutoIncrIdData(db *sql.DB, data string) (int64, error) {
	tblName := "tbl_bench"
	r := mustExec(db, "INSERT INTO "+tblName+"(data) VALUES(?)", data)
	id, err := r.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

/* insert test */
func insertBenchData(db *sql.DB) error {
	for i := 0; i < *N; i++ {
		_, err := insertAutoIncrIdData(db, randSeq(100))
		if err != nil {
			return err
		}
	}
	return nil
}

func concurrentInsertBenchData(db *sql.DB) error {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < *C; i++ {
		go func() {
			for _ = range c {
				_, err := insertAutoIncrIdData(db, randSeq(100))
				if err != nil {
					log.Fatal(err)
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
func readBenchData(db *sql.DB) error {
	for i := 0; i < *N; i++ {
		x := rand.Intn(2000) + 1
		r, err := db.Query("select * from tbl_bench where id = ?", x)
		if err != nil {
			return err
		}
		r.Close()
	}
	return nil
}

func concurrentReadBenchData(db *sql.DB) error {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < *C; i++ {
		go func() {
			for _ = range c {
				x := rand.Intn(2000) + 1
				r, err := db.Query("select * from tbl_bench where id = ?", x)
				if err != nil {
					log.Fatal(err)
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
