package proxy

import (
	"sync"

	"github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/vt/schema"
)

type execTask struct {
	wg   *sync.WaitGroup
	rs   []interface{}
	idx  int
	co   *mysql.SqlConn
	sql  string
	args []interface{}
}

func GetRowCacheType(rowCacheType string) int {
	switch rowCacheType {
	case "RW":
		return schema.CACHE_RW
	case "W":
		return schema.CACHE_W
	default:
		return schema.CACHE_NONE
	}
}
