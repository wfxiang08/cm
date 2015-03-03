package test

import . "gopkg.in/check.v1"

type AutoIncrTestSuit struct {
	createStmts map[string]string
}

var _ = Suite(&AutoIncrTestSuit{
	createStmts: map[string]string{
		"tbl_autoincr": `create table tbl_autoincr (id INT NOT NULL AUTO_INCREMENT, PRIMARY KEY (id), data int)`,
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
}
