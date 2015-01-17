package main

import (
	"testing"

	log "github.com/ngaut/logging"
)

func transaction_testSetup() {
	log.Info("setup suit: transaction_test")
	if isTblExists(`tbl_transaction_test`) {
		mustExec(MysqlDB, `DROP TABLE tbl_transaction_test;`)
	}
	mustExec(MysqlDB, `CREATE TABLE tbl_transaction_test (id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(id), data INT);`)
	reloadConfig()
}

func transaction_testTearDown() {
	log.Info("tear down suit: transaction_test")
	mustExec(MysqlDB, `DROP TABLE tbl_transaction_test;`)
}

func testTransactionCommit(t *testing.T) {
	tx, err := ProxyDB.Begin()
	if err != nil {
		panic(err)
	}

	for i := 2; i < 20; i++ {
		res := mustExec(tx, "insert into "+`tbl_transaction_test`+" (id, data) values (?, ?)", i, i)
		_, err := res.LastInsertId()
		if err != nil {
			t.Error(err)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}

	for i := 2; i < 20; i++ {
		var data int
		mustQueryData(ProxyDB, `tbl_transaction_test`, i, &data)
		if !equal(data, i) {
			t.Error("data != ", i, " return", data)
			return
		}
	}
}

func testTransactionRollback(t *testing.T) {
	tx, err := ProxyDB.Begin()
	if err != nil {
		panic(err)
	}

	for i := 101; i < 120; i++ {
		res := mustExec(tx, "insert into "+`tbl_transaction_test`+" (id, data) values (?, ?)", i, i)
		_, err := res.LastInsertId()
		if err != nil {
			t.Error(err)
			return
		}
	}

	err = tx.Rollback()
	if err != nil {
		t.Fatal(err)
	}

	for i := 101; i < 120; i++ {
		var data int
		queryData(ProxyDB, `tbl_transaction_test`, i, &data)

		if equal(data, i) {
			t.Error("data == ", i, " return", data)
			return
		}
	}
}

func TestAllTransaction_test(t *testing.T) {
	transaction_testSetup()
	defer transaction_testTearDown()
	testTransactionCommit(t)
	testTransactionRollback(t)
}
