package test

import . "gopkg.in/check.v1"

type AutoIncrTestSuit struct {
	createStmts map[string]string
}

var _ = Suite(&AutoIncrTestSuit{
	createStmts: map[string]string{
		"tbl_autoincr": `create table tbl_autoincr(id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY (id), data int)`,
	},
})

func (s *AutoIncrTestSuit) SetUpTest(c *C) {
	dropTables()
	for _, stmt := range s.createStmts {
		mustExec(mysqlDB, stmt)
	}
	reloadConfig()
}

func (s *AutoIncrTestSuit) TearDownTest(c *C) {
	for tblName, _ := range s.createStmts {
		mustExec(mysqlDB, "drop table "+tblName)
	}
}

func (s *AutoIncrTestSuit) Test(c *C) {
	mustExec(proxyDB, "insert into tbl_autoincr(data) values(1024)")
	r := mustExec(proxyDB, "insert into tbl_autoincr(data) values(1025)")
	rowsAffected, err := r.RowsAffected()

	c.Assert(err, Equals, nil)
	c.Assert(rowsAffected, Equals, int64(1))

	rs := mustQuery(proxyDB, "select id from tbl_autoincr where data = 1025")
	defer rs.Close()

	for rs.Next() {
		var id int
		err := rs.Scan(&id)
		c.Assert(err, Equals, nil)
		c.Assert(id, Equals, 2)
	}
}
