package main

import (
	"testing"

	log "github.com/ngaut/logging"
)

func float_testSetup() {
	log.Info("setup suit: float_test")
	if isTblExists(`tbl_float_test`) {
		mustExec(MysqlDB, `DROP TABLE tbl_float_test;`)
	}
	mustExec(MysqlDB, `CREATE TABLE tbl_float_test (id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data DOUBLE);`)
	reloadConfig()
}

func float_testTearDown() {
	log.Info("tear down suit: float_test")
	mustExec(MysqlDB, `DROP TABLE tbl_float_test;`)
}

func float_testTestInsert(t *testing.T) {
	res := mustExec(ProxyDB, "insert into "+`tbl_float_test`+" (id, data) values (?, ?)", 1, 1.5)
	_, err := res.LastInsertId()
	if err != nil {
		t.Error(err)
		return
	}

	var data float64
	mustQueryData(ProxyDB, `tbl_float_test`, 1, &data)
	if !equal(data, 1.5) {
		t.Error("data != ", 1.5, " return", data)
		return
	}

	mustExec(ProxyDB, "delete from tbl_float_test where id = 1")
}

func float_testTestReplace(t *testing.T) {
	res := mustExec(ProxyDB, "replace into "+`tbl_float_test`+" (id, data) values (?, ?)", 1, 1.5)
	_, err := res.LastInsertId()
	if err != nil {
		t.Error(err)
		return
	}

	var data float64
	mustQueryData(ProxyDB, `tbl_float_test`, 1, &data)
	if !equal(data, 1.5) {
		t.Error("data != ", 1.5, " return", data)
		return
	}
}

func float_testTestSelect(t *testing.T) {
	var data float64
	mustQueryData(ProxyDB, `tbl_float_test`, 1, &data)
	if !equal(data, 1.5) {
		t.Error("data != ", 1.5, " return", data)
		return
	}
}

func float_testTestUpdate(t *testing.T) {
	res := mustExec(ProxyDB, "update "+`tbl_float_test`+" set data=? where id=?", 1.6, 1)
	affected, _ := res.RowsAffected()
	if affected != 1 {
		t.Error("affected rows not equals to 1, return", affected)
		return
	}
	var data float64
	mustQueryData(ProxyDB, `tbl_float_test`, 1, &data)
	if !equal(data, 1.6) {
		t.Error("data != ", 1.6, " return", data)
		return
	}
}

func float_testTestDelete(t *testing.T) {
}

func TestAllfloat_test(t *testing.T) {
	float_testSetup()
	defer float_testTearDown()
	float_testTestInsert(t)
	float_testTestReplace(t)
	float_testTestSelect(t)
	float_testTestSelect(t)
	float_testTestUpdate(t)
	float_testTestDelete(t)
}
