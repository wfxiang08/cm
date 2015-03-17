// Copyright 2013, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabletserver

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	stats "github.com/ngaut/vstats"
)

var interval = 5 * time.Second

var mainStringMetrics = map[string]bool{
	"accepting_conns":       false,
	"auth_cmds":             false,
	"auth_errors":           false,
	"bytes_read":            false,
	"bytes_written":         false,
	"bytes":                 false,
	"cas_badval":            false,
	"cas_hits":              false,
	"cas_misses":            false,
	"cmd_flush":             false,
	"cmd_get":               false,
	"cmd_set":               false,
	"cmd_touch":             false,
	"conn_yields":           false,
	"connection_structures": false,
	"curr_connections":      false,
	"curr_items":            false,
	"decr_hits":             false,
	"decr_misses":           false,
	"delete_hits":           false,
	"delete_misses":         false,
	"evicted_unfetched":     false,
	"evictions":             false,
	"expired_unfetched":     false,
	"get_hits":              false,
	"get_misses":            false,
	"hash_bytes":            false,
	"hash_is_expanding":     false,
	"hash_power_level":      false,
	"incr_hits":             false,
	"incr_misses":           false,
	"libevent":              true,
	"limit_maxbytes":        false,
	"listen_disabled_num":   false,
	"pid":               false,
	"pointer_size":      false,
	"reclaimed":         false,
	"reserved_fds":      false,
	"rusage_system":     true,
	"rusage_user":       true,
	"threads":           false,
	"time":              false,
	"total_connections": false,
	"total_items":       false,
	"touch_hits":        false,
	"touch_misses":      false,
	"uptime":            false,
	"version":           true,
}

var slabsSingleMetrics = map[string]bool{
	"active_slabs":    true,
	"cas_badval":      false,
	"cas_hits":        false,
	"chunk_size":      false,
	"chunks_per_page": false,
	"cmd_set":         false,
	"decr_hits":       false,
	"delete_hits":     false,
	"free_chunks_end": false,
	"free_chunks":     false,
	"get_hits":        false,
	"incr_hits":       false,
	"mem_requested":   false,
	"total_chunks":    false,
	"total_malloced":  true,
	"total_pages":     false,
	"touch_hits":      false,
	"used_chunks":     false,
}

var itemsMetrics = []string{
	"age",
	"evicted",
	"evicted_nonzero",
	"evicted_time",
	"evicted_unfetched",
	"expired_unfetched",
	"number",
	"outofmemory",
	"reclaimed",
	"tailrepairs",
}

var internalErrors *stats.Counters

func formatKey(key string) string {
	key = regexp.MustCompile("^[a-z]").ReplaceAllStringFunc(key, func(item string) string {
		return strings.ToUpper(item)
	})
	key = regexp.MustCompile("_[a-z]").ReplaceAllStringFunc(key, func(item string) string {
		return strings.ToUpper(item[1:])
	})
	return key
}

// parseSlabKey splits a slab key into the subkey and slab id:
// "1:chunk_size" -> "chunk_size", 1
func parseSlabKey(key string) (subkey string, slabid string, err error) {
	tokens := strings.Split(key, ":")
	if len(tokens) != 2 {
		return "", "", fmt.Errorf("invalid slab key: %v", key)
	}
	return tokens[1], tokens[0], nil
}

// parseItemKey splits an item key into the subkey and slab id:
// "items:1:number" -> "number", 1
func parseItemKey(key string) (subkey string, slabid string, err error) {
	tokens := strings.Split(key, ":")
	if len(tokens) != 3 {
		return "", "", fmt.Errorf("invalid slab key: %v", key)
	}
	return tokens[2], tokens[1], nil
}

func copyMap(src map[string]int64) map[string]int64 {
	if src == nil {
		return nil
	}
	dst := make(map[string]int64, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
