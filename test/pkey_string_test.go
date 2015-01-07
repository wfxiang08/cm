package main

import (
	"testing"

	log "github.com/ngaut/logging"
)

func pkey_string_testSetup() {
	log.Info("setup suit: pkey_string_test")
	if isTblExists(`tbl_pkey_string_test`) {
		mustExec(MysqlDB, `DROP TABLE tbl_pkey_string_test;`)
	}
	mustExec(MysqlDB, `CREATE TABLE tbl_pkey_string_test (id VARCHAR(20), UNIQUE KEY(id), data VARCHAR(20));`)
	reloadConfig()
}

func pkey_string_testTearDown() {
	log.Info("tear down suit: pkey_string_test")
	mustExec(MysqlDB, `DROP TABLE tbl_pkey_string_test;`)
}

func pkey_string_testTestInsert(t *testing.T) {
	mustExec(ProxyDB, "insert into "+`tbl_pkey_string_test`+" (id, data) values (?, ?)", `1`, "hello world")

	var data string
	mustQueryData(ProxyDB, `tbl_pkey_string_test`, `1`, &data)
	if !equal(data, "hello world") {
		t.Error("data != ", "hello world", " return", data)
		return
	}
}

func pkey_string_testTestSelect(t *testing.T) {
	var data string
	mustQueryData(ProxyDB, `tbl_pkey_string_test`, `1`, &data)
	if !equal(data, "hello world") {
		t.Error("data != ", "hello world", " return", data)
		return
	}
}

func pkey_string_testTestUpdate(t *testing.T) {
	res := mustExec(ProxyDB, "update "+`tbl_pkey_string_test`+" set data=? where id=?", "hello world!", `1`)
	affected, _ := res.RowsAffected()
	if affected != 1 {
		t.Error("affected rows not equals to 1, return", affected)
		return
	}
	var data string
	mustQueryData(ProxyDB, `tbl_pkey_string_test`, `1`, &data)
	if !equal(data, "hello world!") {
		t.Error("data != ", "hello world!", " return", data)
		return
	}
}

func pkey_string_testTestDelete(t *testing.T) {
}

func TestAllpkey_string_test(t *testing.T) {
	pkey_string_testSetup()
	defer pkey_string_testTearDown()
	pkey_string_testTestInsert(t)
	pkey_string_testTestSelect(t)
	pkey_string_testTestSelect(t)
	pkey_string_testTestUpdate(t)
	pkey_string_testTestDelete(t)
}
