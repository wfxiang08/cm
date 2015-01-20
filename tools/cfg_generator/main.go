package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	log "github.com/ngaut/logging"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlHost = flag.String("h", "127.0.0.1", "mysql host")
	mysqlPort = flag.Int("p", 3306, "mysql port")
	mysqlUser = flag.String("u", "root", "user")
	mysqlPwd  = flag.String("P", "", "password")
	dbName    = flag.String("db", "test", "db name")
	cacheType = flag.String("cache", "", "row cache: R RW")
)

const (
	tmpl = `
{
    "addr": "127.0.0.1:4000",
    "log_level": "warning",
    "shards": [
        {
            "down_after_noalive": 0,
            "idle_conns": 100,
            "master": "127.0.0.1:3306",
            "name": "shard1",
            "password": "",
            "rw_split": false,
            "slave": "",
            "user": "root"
        }
    ],
	"rowcache_conf":{
		"binary":"/usr/bin/memcached",
		"mem":128,
		"socket":"",
		"port":11222,
		"connections":1024,
		"threads":-1,
		"lock_paged":false
    },
    "password": "",
    "schemas": [
        {
            "db": "{{ .DbName }}",
            "shards": [
                "shard1"
            ],
            "rules": {
                "default": "shard1",
                "shard": {{ .Shards }}
            }
        }
    ],
    "user": "root"
}
`
)

type ShardInfo struct {
	Table        string `json:"table"`
	RowCacheType string `json:"row_cache_type,omitempty"`
	Key          string `json:"key"`
}

func NewDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	return db, nil
}

func NewMysqlDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *mysqlUser, *mysqlPwd, *mysqlHost, *mysqlPort, *dbName)
	return NewDb(dsn)
}

func main() {
	flag.Parse()
	// show tables
	log.SetOutput(os.Stderr)

	db, err := NewMysqlDb()
	if err != nil {
		log.Fatal(err)
	}

	r, err := db.Query("SHOW TABLES")
	var tbls []string

	for r.Next() {
		var tblName string
		r.Scan(&tblName)
		tbls = append(tbls, tblName)
	}
	r.Close()

	var shards []ShardInfo
	for _, tbl := range tbls {
		r, err := db.Query("SHOW COLUMNS FROM " + tbl)
		if err != nil {
			log.Fatal(err)
		}

		var priKey string
		var uniqueKey string
		var mulKeys []string
		for r.Next() {
			values := make([]sql.RawBytes, 6)
			scanArgs := make([]interface{}, 6)
			for i := range values {
				scanArgs[i] = &values[i]
			}
			err := r.Scan(scanArgs...)
			if err != nil {
				log.Fatal(err)
			}
			if string(values[3]) == "UNI" {
				uniqueKey = string(values[0])
			} else if string(values[3]) == "MUL" {
				mulKeys = append(mulKeys, string(values[0]))
			} else if string(values[3]) == "PRI" {
				priKey = string(values[0])
			}
		}
		mulKey := strings.Join(mulKeys, ",")

		var key string
		if len(uniqueKey) > 0 {
			key = uniqueKey
		} else if len(mulKey) > 0 {
			key = mulKey
		} else if len(priKey) > 0 {
			key = priKey
		} else {
			log.Warning("illgial table", tbl)
			r.Close()
			break
		}
		info := ShardInfo{
			Table: tbl,
			Key:   key,
		}
		if len(*cacheType) > 0 {
			info.RowCacheType = *cacheType
		}
		shards = append(shards, info)
		r.Close()
	}

	shardsJson, _ := json.Marshal(shards)

	buf := bytes.NewBuffer(nil)

	t := template.Must(template.New("cfg_tmpl").Parse(tmpl))
	t.Execute(buf, &struct {
		DbName string
		Shards string
	}{
		*dbName,
		string(shardsJson),
	})

	// prettyify
	var ret interface{}
	err = json.Unmarshal(buf.Bytes(), &ret)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(ret, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
}
