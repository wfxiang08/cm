package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/c4pt0r/mysql"
	"github.com/kyokomi/emoji"
)

func createCharsetTestTbls(db *sql.DB) {
	if b, err := isTblExists(db, "utf8mb4_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE utf8mb4_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data VARCHAR(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "gb2312_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE gb2312_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data VARCHAR(1024) CHARACTER SET gb2312 COLLATE gb2312_chinese_ci)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "gbk_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE gbk_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data VARCHAR(1024) CHARACTER SET gbk COLLATE gbk_chinese_ci)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "utf8_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE utf8_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data VARCHAR(1024) CHARACTER SET utf8 COLLATE utf8_general_ci)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "utf16_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE utf16_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data VARCHAR(1024) CHARACTER SET utf16 COLLATE utf16_general_ci)`)
	} else if err != nil {
		log.Fatal(err)
	}

	if b, err := isTblExists(db, "utf32_test"); !b && err == nil {
		mustExec(db, `CREATE TABLE utf32_test(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data VARCHAR(1024) CHARACTER SET utf32 COLLATE utf32_general_ci)`)
	} else if err != nil {
		log.Fatal(err)
	}

}

func dropCharsetTestTbls(db *sql.DB) {
	tbls := []string{
		"utf8mb4_test",
		"gb2312_test",
		"gbk_test",
		"utf16_test",
		"utf32_test",
	}

	for _, t := range tbls {
		if b, _ := isTblExists(db, t); b {
			mustExec(db, "DROP TABLE "+t)
		}
	}
}

func utf8mb4Test(db *sql.DB) error {
	s := emoji.Sprint("I like a :pizza: and :sushi:!!")
	var ret, cret []byte
	err := insertDataAndQueryBack(db, "utf8mb4_test", s, &ret, &cret)
	if err != nil {
		return err
	}

	if string(ret) != s {
		return fmt.Errorf("emoji test failed %v != %v", ret, []byte(s))
	}

	if string(ret) != string(cret) {
		return fmt.Errorf("emoji cache test failed %v != %v", ret, cret)
	}
	return nil
}

func gb2312Test(db *sql.DB) error {
	s := "\xb2\xe2\xca\xd4"
	var ret, cret []byte
	err := insertDataAndQueryBack(db, "gb2312_test", s, &ret, &cret)
	if err != nil {
		return err
	}

	if string(ret) != s {
		return fmt.Errorf("gb2312 test failed %v != %v", ret, []byte(s))
	}

	if string(ret) != string(cret) {
		return fmt.Errorf("gb2312 cache test failed %v != %v", ret, cret)
	}
	return nil
}

func gbkTest(db *sql.DB) error {
	s := "\xb2\xe2\xca\xd4"
	var ret, cret []byte
	err := insertDataAndQueryBack(db, "gbk_test", s, &ret, &cret)
	if err != nil {
		return err
	}

	if string(ret) != s {
		return fmt.Errorf("gb2312 test failed %v != %v", ret, []byte(s))
	}

	if string(ret) != string(cret) {
		return fmt.Errorf("gb2312 cache test failed %v != %v", ret, cret)
	}
	return nil
}
