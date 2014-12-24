package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	_ "github.com/c4pt0r/mysql"
	"github.com/kyokomi/emoji"
)

func createTypeTestTbls(db *sql.DB) {
	if b, err := isTblExists(db, "int_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE int_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data INT)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "double_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE double_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data DOUBLE)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "varchar_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE varchar_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data VARCHAR(1024))`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "text_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE text_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data TEXT)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "blob_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE blob_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data BLOB)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "datetime_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE datetime_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data DATETIME)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "date_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE date_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data DATE)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "multi_primary_key_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE multi_primary_key_test(id1 INT, id2 INT, PRIMARY KEY(id1, id2), data INT)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "emoji_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE emoji_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data VARCHAR(1024))`)
	} else if err != nil {
		log.Fatal(err)
	}

}

func dropTypeTestTbls(db *sql.DB) {
	tbls := []string{
		"int_test",
		"double_test",
		"varchar_test",
		"text_test",
		"blob_test",
		"datetime_test",
		"date_test",
	}

	for _, t := range tbls {
		if b, _ := isTblExists(db, t); b {
			mustExec(db, "DROP TABLE "+t)
		}
	}
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

func multiPKeyTest(db *sql.DB) error {
	x := rand.Intn(1024)
	id1 := rand.Intn(1024)
	id2 := rand.Intn(1024)
	tblName := "multi_primary_key_test"

	mustExec(db, "INSERT INTO multi_primary_key_test(id1, id2, data) VALUES(?, ?, ?)", id1, id2, x)

	res, err := db.Query("SELECT data FROM "+tblName+" WHERE id1 = ? and id2 = ?", id1, id2)
	if err != nil {
		return err
	}
	defer res.Close()

	for res.Next() {
		var ret int
		res.Scan(&ret)
		if err != nil {
			return err
		}
		if ret != x {
			return fmt.Errorf("multi pkey test failed %d != %d", x, ret)
		}
	}

	cachedRes, err := db.Query("SELECT data FROM "+tblName+" WHERE id1 = ? and id2 = ?", id1, id2)
	if err != nil {
		return err
	}
	defer cachedRes.Close()

	for cachedRes.Next() {
		var ret int
		cachedRes.Scan(&ret)
		if err != nil {
			return err
		}
		if ret != x {
			return fmt.Errorf("multi pkey test cache failed %d != %d", ret, x)
		}
	}
	return nil
}

func intTest(db *sql.DB) error {
	x := rand.Intn(1024)
	var ret, cret int
	err := insertDataAndQueryBack(db, "int_test", x, &ret, &cret)
	if err != nil {
		return err
	}
	if ret != x {
		return fmt.Errorf("int test failed %d != %d", ret, x)
	}

	if ret != cret {
		return fmt.Errorf("int test cache failed %d != %d", ret, cret)
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

	if math.Abs(ret-cret) > 1e-7 {
		return fmt.Errorf("double test cache failed %v != %v", ret, cret)
	}

	return nil
}

func emojiTest(db *sql.DB) error {
	s := emoji.Sprint("I like a :pizza: and :sushi:!!")
	var ret, cret []byte
	err := insertDataAndQueryBack(db, "emoji_test", s, &ret, &cret)
	if err != nil {
		return err
	}

	if len(ret) != len(s) {
		return fmt.Errorf("emoji test failed %v != %v", ret, []byte(s))
	}

	if len(ret) != len(cret) {
		return fmt.Errorf("emoji cache test failed %v != %v", ret, cret)
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

	if ret != cret {
		return fmt.Errorf("var char cache test failed %v != %v", ret, cret)
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

	if ret != cret {
		return fmt.Errorf("text test cache failed %v != %v", ret, cret)
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
	s3 := fmt.Sprintf("%x", cret)
	if s1 != s2 {
		return fmt.Errorf("blob test failed %v != %v", s1, s2)
	}

	if s1 != s3 {
		return fmt.Errorf("blob test cache failed %v != %v", s1, s3)
	}

	return nil
}

func datetimeTest(db *sql.DB) error {
	d := time.Now()
	ds := d.Format("2006-01-02 15:04:05")
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

	if string(cret) != string(ret) {
		return fmt.Errorf("datetime test cache failed %v != %v", string(ret), string(cret))
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

	if string(ret) != string(cret) {
		return fmt.Errorf("date test cache failed %v != %v", string(ret), string(cret))
	}

	return nil
}
