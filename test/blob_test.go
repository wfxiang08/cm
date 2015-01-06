package main

import (
	"testing"

	log "github.com/ngaut/logging"
)

func blob_testSetup() {
	log.Info("setup suit: blob_test")
	if isTblExists(`tbl_blob_test`) {
		mustExec(MysqlDB, `DROP TABLE tbl_blob_test;`)
	}
	mustExec(MysqlDB, `CREATE TABLE tbl_blob_test (id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data BLOB);`)
	reloadConfig()
}

func blob_testTearDown() {
	log.Info("tear down suit: blob_test")
	mustExec(MysqlDB, `DROP TABLE tbl_blob_test;`)
}

func blob_testTestInsert(t *testing.T) {
	res := mustExec(ProxyDB, "insert into "+`tbl_blob_test`+" (id, data) values (?, ?)", 1, []byte("\xff"))
	id, err := res.LastInsertId()
	if err != nil {
		t.Error(err)
		return
	}
	if id != 1 {
		t.Error("id not equals to 1, return", id)
		return
	}

	var data []byte
	mustQueryData(ProxyDB, `tbl_blob_test`, id, &data)
	if !equal(data, []byte("\xff")) {
		t.Error("data != ", []byte("\xff"), " return", data)
		return
	}
}

func blob_testTestSelect(t *testing.T) {
	var data []byte
	mustQueryData(ProxyDB, `tbl_blob_test`, 1, &data)
	if !equal(data, []byte("\xff")) {
		t.Error("data != ", []byte("\xff"), " return", data)
		return
	}
}

func blob_testTestUpdate(t *testing.T) {
	res := mustExec(ProxyDB, "update "+`tbl_blob_test`+" set data=? where id=?", []byte("\xfe"), 1)
	affected, _ := res.RowsAffected()
	if affected != 1 {
		t.Error("affected rows not equals to 1, return", affected)
		return
	}
	var data []byte
	mustQueryData(ProxyDB, `tbl_blob_test`, 1, &data)
	if !equal(data, []byte("\xfe")) {
		t.Error("data != ", []byte("\xfe"), " return", data)
		return
	}
}

func blob_testTestDelete(t *testing.T) {
}

func TestAllblob_test(t *testing.T) {
	blob_testSetup()
	defer blob_testTearDown()
	blob_testTestInsert(t)
	blob_testTestSelect(t)
	blob_testTestSelect(t)
	blob_testTestUpdate(t)
	blob_testTestDelete(t)
}
