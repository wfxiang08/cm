package main

import (
	"testing"

	log "github.com/ngaut/logging"
)

func multi_pkey_testSetup() {
	log.Info("setup suit: multi_pkey_test")
	if isTblExists(`tbl_multi_pkey_test`) {
		mustExec(MysqlDB, `DROP TABLE tbl_multi_pkey_test;`)
	}
	mustExec(MysqlDB, `CREATE TABLE tbl_multi_pkey_test (id1 VARCHAR(20), id2 VARCHAR(20), UNIQUE KEY(id1, id2), data INT);`)
	reloadConfig()
}

func multi_pkey_testTearDown() {
	log.Info("tear down suit: multi_pkey_test")
	mustExec(MysqlDB, `DROP TABLE tbl_multi_pkey_test;`)
}

func multi_pkey_testTestInsert(t *testing.T) {
	mustExec(ProxyDB, "insert into "+`tbl_multi_pkey_test`+" (id1, id2, data) values (?, ?, ?)", `id1`, `id2`, 100)
	var data int
	mustQueryDataWithMultiId(ProxyDB, `tbl_multi_pkey_test`, `id1`, `id2`, &data)
	if !equal(data, 100) {
		t.Error("data != ", 100, " return", data)
		return
	}
}

func multi_pkey_testTestSelect(t *testing.T) {
	var data int
	mustQueryDataWithMultiId(ProxyDB, `tbl_multi_pkey_test`, `id1`, `id2`, &data)
	if !equal(data, 100) {
		t.Error("data != ", 100, " return", data)
		return
	}
}

func multi_pkey_testTestUpdate(t *testing.T) {
	res := mustExec(ProxyDB, "update "+`tbl_multi_pkey_test`+" set data=? where id1=? and id2=?", 200, `id1`, `id2`)
	affected, _ := res.RowsAffected()
	if affected != 1 {
		t.Error("affected rows not equals to 1, return", affected)
		return
	}
	var data int
	mustQueryDataWithMultiId(ProxyDB, `tbl_multi_pkey_test`, `id1`, `id2`, &data)
	if !equal(data, 200) {
		t.Error("data != ", 200, " return", data)
		return
	}
}

func multi_pkey_testTestDelete(t *testing.T) {
}

func TestAllmulti_pkey_test(t *testing.T) {
	multi_pkey_testSetup()
	defer multi_pkey_testTearDown()
	multi_pkey_testTestInsert(t)
	multi_pkey_testTestSelect(t)
	multi_pkey_testTestSelect(t)
	multi_pkey_testTestUpdate(t)
	multi_pkey_testTestDelete(t)
}
