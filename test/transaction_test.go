package test

import . "gopkg.in/check.v1"

type TransTestSuit struct {
	createStmts map[string]string
}

var _ = Suite(&TransTestSuit{
	createStmts: map[string]string{
		"tbl_trans": `create table tbl_trans (id int, data varchar(255))`,
	},
})

func (s *TransTestSuit) SetUpTest(c *C) {
	dropTables()
	for _, stmt := range s.createStmts {
		mustExec(mysqlDB, stmt)
	}
	reloadConfig()
}

func (s *TransTestSuit) TearDownTest(c *C) {
	for tblName, _ := range s.createStmts {
		mustExec(mysqlDB, "drop table "+tblName)
	}
}

func (s *TransTestSuit) TestTrans(c *C) {
	// commit
	tx, err := proxyDB.Begin()
	for i := 0; i < 1000; i++ {
		tx.Exec("insert into tbl_trans values(?, 'a')", i)
	}
	err = tx.Commit()
	c.Assert(err, Equals, nil)

	rs := mustQuery(proxyDB, "select count(*) from tbl_trans")
	defer rs.Close()

	for rs.Next() {
		var cnt int
		err := rs.Scan(&cnt)
		c.Assert(err, Equals, nil)
		c.Assert(cnt, Equals, 1000)
	}

	// rollback
	tx, err = proxyDB.Begin()
	for i := 0; i < 1000; i++ {
		tx.Exec("insert into tbl_trans values(?, 'a')", i)
	}
	err = tx.Rollback()
	c.Assert(err, Equals, nil)

	for rs.Next() {
		var cnt int
		err := rs.Scan(&cnt)
		c.Assert(err, Equals, nil)
		c.Assert(cnt, Equals, 1000)
	}
}
