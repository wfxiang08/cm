package main

import (
	"testing"

	log "github.com/ngaut/logging"
)

func uint_testSetup() {
	log.Info("setup suit: uint_test")
	if isTblExists(`tbl_uint_test`) {
		mustExec(MysqlDB, `DROP TABLE tbl_uint_test;`)
	}
	mustExec(MysqlDB, `CREATE TABLE tbl_uint_test (id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data INT UNSIGNED);`)
	reloadConfig()
}

func uint_testTearDown() {
	log.Info("tear down suit: uint_test")
	mustExec(MysqlDB, `DROP TABLE tbl_uint_test;`)
}

func uint_testTestInsert(t *testing.T) {
	res := mustExec(ProxyDB, "insert into "+`tbl_uint_test`+" (id, data) values (?, ?)", 1, uint64(100))
	_, err := res.LastInsertId()
	if err != nil {
		t.Error(err)
		return
	}

	var data uint64
	mustQueryData(ProxyDB, `tbl_uint_test`, 1, &data)
	if !equal(data, uint64(100)) {
		t.Error("data != ", uint64(100), " return", data)
		return
	}

	mustExec(ProxyDB, "delete from tbl_uint_test where id = 1")
}

func uint_testTestReplace(t *testing.T) {
	res := mustExec(ProxyDB, "replace into "+`tbl_uint_test`+" (id, data) values (?, ?)", 1, uint64(100))
	_, err := res.LastInsertId()
	if err != nil {
		t.Error(err)
		return
	}

	var data uint64
	mustQueryData(ProxyDB, `tbl_uint_test`, 1, &data)
	if !equal(data, uint64(100)) {
		t.Error("data != ", uint64(100), " return", data)
		return
	}
}

func uint_testTestReplace2(t *testing.T) {
	res := mustExec(ProxyDB, "replace into "+`tbl_uint_test`+" (id, data) values (?, ?)", 1, uint64(200))
	_, err := res.LastInsertId()
	if err != nil {
		t.Error(err)
		return
	}

	var data uint64
	mustQueryData(ProxyDB, `tbl_uint_test`, 1, &data)
	if !equal(data, uint64(200)) {
		t.Error("data != ", uint64(200), " return", data)
		return
	}
}

func uint_testTestSelect(t *testing.T) {
	var data uint64
	mustQueryData(ProxyDB, `tbl_uint_test`, 1, &data)
	if !equal(data, uint64(100)) {
		t.Error("data != ", uint64(100), " return", data)
		return
	}
}

func uint_testTestUpdate(t *testing.T) {
	res := mustExec(ProxyDB, "update "+`tbl_uint_test`+" set data=? where id=?", uint64(200), 1)
	affected, _ := res.RowsAffected()
	if affected != 1 {
		t.Error("affected rows not equals to 1, return", affected)
		return
	}
	var data uint64
	mustQueryData(ProxyDB, `tbl_uint_test`, 1, &data)
	if !equal(data, uint64(200)) {
		t.Error("data != ", uint64(200), " return", data)
		return
	}
}

func uint_testTestDelete(t *testing.T) {
}

func TestAlluint_test(t *testing.T) {
	uint_testSetup()
	defer uint_testTearDown()
	uint_testTestInsert(t)
	uint_testTestReplace(t)
	uint_testTestReplace2(t)
	uint_testTestReplace(t)
	uint_testTestSelect(t)
	uint_testTestSelect(t)
	uint_testTestUpdate(t)
	uint_testTestDelete(t)
}
