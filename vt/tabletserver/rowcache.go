// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabletserver

import (
	"encoding/binary"
	"strconv"
	"time"

	log "github.com/ngaut/logging"

	"github.com/wandoulabs/cm/mysql"
	"github.com/wandoulabs/cm/vt/schema"
	"github.com/youtube/vitess/go/stats"
)

var cacheStats = stats.NewTimings("Rowcache")

var pack = binary.BigEndian

const (
	RC_DELETED = 1

	// MAX_KEY_LEN is a value less than memcache's limit of 250.
	MAX_KEY_LEN = 200

	// MAX_DATA_LEN prevents large rows from being inserted in rowcache.
	MAX_DATA_LEN = 8000
)

type RowCache struct {
	tableInfo *TableInfo
	prefix    string
	cachePool *CachePool
}

type RCResult struct {
	Row mysql.RowValue
	Cas uint64
}

func NewRowCache(tableInfo *TableInfo, cachePool *CachePool) *RowCache {
	prefix := strconv.FormatInt(cachePool.maxPrefix.Add(1), 36) + "."
	return &RowCache{tableInfo, prefix, cachePool}
}

func (rc *RowCache) Get(keys []string, tcs []schema.TableColumn) (results map[string]RCResult) {
	mkeys := make([]string, 0, len(keys))
	for _, key := range keys {
		if len(key) > MAX_KEY_LEN {
			continue
		}
		mkeys = append(mkeys, rc.prefix+key)
	}
	prefixlen := len(rc.prefix)
	conn := rc.cachePool.Get(0)
	// This is not the same as defer rc.cachePool.Put(conn)
	defer func() { rc.cachePool.Put(conn) }()

	defer cacheStats.Record("Exec", time.Now())
	mcresults, err := conn.Gets(mkeys...)
	if err != nil {
		conn.Close()
		conn = nil
		log.Fatalf("%s", err)
	}
	results = make(map[string]RCResult, len(mkeys))
	for _, mcresult := range mcresults {
		if mcresult.Flags == RC_DELETED {
			// The row was recently invalidated.
			// If the caller reads the row from db, they can update it
			// back as long as it's not updated again.
			results[mcresult.Key[prefixlen:]] = RCResult{Cas: mcresult.Cas}
			continue
		}
		row := rc.decodeRow(mcresult.Value, tcs)
		if row == nil {
			log.Fatalf("Corrupt data for %s", mcresult.Key)
		}
		results[mcresult.Key[prefixlen:]] = RCResult{Row: row, Cas: mcresult.Cas}
	}
	return
}

func (rc *RowCache) Set(key string, row []byte, cas uint64) {
	if len(key) > MAX_KEY_LEN {
		return
	}

	conn := rc.cachePool.Get(0)
	defer func() { rc.cachePool.Put(conn) }()
	mkey := rc.prefix + key

	var err error
	if cas == 0 {
		// Either caller didn't find the value at all
		// or they didn't look for it in the first place.
		_, err = conn.Add(mkey, 0, 0, row)
	} else {
		// Caller is trying to update a row that recently changed.
		_, err = conn.Cas(mkey, 0, 0, row, cas)
	}
	if err != nil {
		conn.Close()
		conn = nil
		log.Fatalf("%s", err)
	}
}

func (rc *RowCache) Delete(key string) {
	if len(key) > MAX_KEY_LEN {
		return
	}
	conn := rc.cachePool.Get(0)
	defer func() { rc.cachePool.Put(conn) }()
	mkey := rc.prefix + key

	_, err := conn.Set(mkey, RC_DELETED, rc.cachePool.DeleteExpiry, nil)
	if err != nil {
		conn.Close()
		conn = nil
		log.Fatalf("%s", err)
	}
}

func (rc *RowCache) decodeRow(b []byte, tcs []schema.TableColumn) mysql.RowValue {
	fs := make([]*mysql.Field, 0, len(tcs))

	for _, tc := range tcs {
		f := &mysql.Field{
			Type: uint8(tc.Category),
		}
		fs = append(fs, f)
	}

	row, err := mysql.RowData(b).ParseText(fs)
	if err != nil {
		log.Error(err)
	}

	return row
}
