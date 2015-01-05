package main

import (
	"testing"

	log "github.com/ngaut/logging"
)

func int_testSetup() {
    log.Info("setup suit: int_test")
    if isTblExists(`tbl_int_test`) {
        mustExec(MysqlDB, `DROP TABLE tbl_int_test;`)
	}
    mustExec(MysqlDB,`CREATE TABLE tbl_int_test (id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data INT);`)
	reloadConfig()
}

func int_testTearDown() {
    log.Info("tear down suit: int_test")
    mustExec(MysqlDB, `DROP TABLE tbl_int_test;`)
}

func int_testTestInsert(t *testing.T) {
    res := mustExec(ProxyDB, "insert into "+`tbl_int_test`+" (id, data) values (?, ?)", 1, 100)
	id, err := res.LastInsertId()
	if err != nil {
		t.Error(err)
		return
	}
    if id != 1 {
        t.Error("id not equals to 1, return", id)
		return
	}

    var data int
	mustQueryData(ProxyDB, `tbl_int_test`, id, &data)
    if !equal(data, 100) {
        t.Error("data != ", 100, " return", data)
		return
	}
}

func int_testTestSelect(t *testing.T) {
    var data int
    mustQueryData(ProxyDB, `tbl_int_test`, 1, &data)
	if !equal(data, 100) {
		t.Error("data != ", 100, " return", data)
		return
	}
}

func int_testTestUpdate(t *testing.T) {
    res := mustExec(ProxyDB, "update "+`tbl_int_test`+" set data=? where id=?", 200, 1)
	affected, _ := res.RowsAffected()
	if affected != 1 {
		t.Error("affected rows not equals to 1, return", affected)
		return
	}
    var data int
    mustQueryData(ProxyDB, `tbl_int_test`, 1, &data)
	if !equal(data, 200) {
		t.Error("data != ", 200 ," return", data)
		return
	}
}

func int_testTestDelete(t *testing.T) {
}

func TestAllint_test(t *testing.T) {
    int_testSetup()
    defer int_testTearDown()
    int_testTestInsert(t)
	int_testTestSelect(t)
	int_testTestSelect(t)
	int_testTestUpdate(t)
	int_testTestDelete(t)
}
