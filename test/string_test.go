package main

import (
	"testing"

	log "github.com/ngaut/logging"
)

func string_testSetup() {
	log.Info("setup suit: string_test")
	if isTblExists(`tbl_string_test`) {
		mustExec(MysqlDB, `DROP TABLE tbl_string_test;`)
	}
	mustExec(MysqlDB, `CREATE TABLE tbl_string_test (id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data TEXT);`)
	reloadConfig()
}

func string_testTearDown() {
	log.Info("tear down suit: string_test")
	mustExec(MysqlDB, `DROP TABLE tbl_string_test;`)
}

func string_testTestInsert(t *testing.T) {
	res := mustExec(ProxyDB, "insert into "+`tbl_string_test`+" (id, data) values (?, ?)", 1, "hello world")
	_, err := res.LastInsertId()
	if err != nil {
		t.Error(err)
		return
	}

	var data string
	mustQueryData(ProxyDB, `tbl_string_test`, 1, &data)
	if !equal(data, "hello world") {
		t.Error("data != ", "hello world", " return", data)
		return
	}

	mustExec(ProxyDB, "delete from tbl_string_test where id = 1")
}

func string_testTestReplace(t *testing.T) {
	res := mustExec(ProxyDB, "replace into "+`tbl_string_test`+" (id, data) values (?, ?)", 1, "hello world")
	_, err := res.LastInsertId()
	if err != nil {
		t.Error(err)
		return
	}

	var data string
	mustQueryData(ProxyDB, `tbl_string_test`, 1, &data)
	if !equal(data, "hello world") {
		t.Error("data != ", "hello world", " return", data)
		return
	}
}

func string_testTestSelect(t *testing.T) {
	var data string
	mustQueryData(ProxyDB, `tbl_string_test`, 1, &data)
	if !equal(data, "hello world") {
		t.Error("data != ", "hello world", " return", data)
		return
	}
}

func string_testTestUpdate(t *testing.T) {
	res := mustExec(ProxyDB, "update "+`tbl_string_test`+" set data=? where id=?", "hello world!", 1)
	affected, _ := res.RowsAffected()
	if affected != 1 {
		t.Error("affected rows not equals to 1, return", affected)
		return
	}
	var data string
	mustQueryData(ProxyDB, `tbl_string_test`, 1, &data)
	if !equal(data, "hello world!") {
		t.Error("data != ", "hello world!", " return", data)
		return
	}
}

func string_testTestDelete(t *testing.T) {
}

func TestAllstring_test(t *testing.T) {
	string_testSetup()
	defer string_testTearDown()
	string_testTestInsert(t)
	string_testTestReplace(t)
	string_testTestSelect(t)
	string_testTestSelect(t)
	string_testTestUpdate(t)
	string_testTestDelete(t)
}
