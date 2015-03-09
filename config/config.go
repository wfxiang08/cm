package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

type SchemaConfig struct {
	DB           string       `json:"db" toml:"db"`
	ShardIds     []string     `json:"shard_ids" toml:"shard_ids"`
	RouterConifg RouterConfig `json:"router" toml:"router"`
	CacheSize    int          `json:"cache_size,string" toml:"cache_size"` //m
}

type RouterConfig struct {
	Default   []string    `json:"default_shards" toml:"default_shards"`
	TableRule []TableRule `json:"table_rules" toml:"table_rules"`
}

type TableRule struct {
	Table        string   `json:"table" toml:"table"`
	ShardingKey  string   `json:"key" toml:"key"`
	RowCacheType string   `json:"row_cache_type" toml:"row_cache_type"`
	MapToShards  []string `json:"map_to_shards" toml:"map_to_shards"` //shard ids
}

type ShardConfig struct {
	Id       string `json:"id" toml:"id"`
	User     string `json:"user" toml:"user"`
	Password string `json:"password" toml:"password"`

	Master string `json:"master" toml:"master"`
	Slave  string `json:"slave" toml:"slave"`
}

type Config struct {
	Addr         string         `json:"addr" toml:"addr"`
	User         string         `json:"user" toml:"user"`
	Password     string         `json:"password" toml:"password"`
	LogLevel     string         `json:"log_level" toml:"log_level"`
	SkipAuth     bool           `json:"skip_auth" toml:"skip_auth"`
	Shards       []ShardConfig  `json:"shards" toml:"shards"`
	Schemas      []SchemaConfig `json:"schemas" toml:"schemas"`
	RowCacheConf RowCacheConfig `json:"rowcache_conf" toml:"rowcache_conf"`
}

type RowCacheConfig struct {
	Binary      string `json:"binary" toml:"binary"`
	Memory      int    `json:"mem" toml:"mem"`
	Socket      string `json:"socket" toml:"socket"`
	TcpPort     int    `json:"port" toml:"port"`
	Connections int    `json:"connections" toml:"connections"`
	Threads     int    `json:"threads" toml:"threads"`
	LockPaged   bool   `json:"lock_paged" toml:"lock_paged"`
}

func (c *RowCacheConfig) GetSubprocessFlags() []string {
	cmd := []string{}
	if c.Binary == "" {
		return cmd
	}
	cmd = append(cmd, c.Binary)
	if c.Memory > 0 {
		// memory is given in bytes and rowcache expects in MBs
		cmd = append(cmd, "-m", strconv.Itoa(c.Memory))
	}
	if c.Socket != "" {
		cmd = append(cmd, "-s", c.Socket)
	}
	if c.TcpPort > 0 {
		cmd = append(cmd, "-p", strconv.Itoa(c.TcpPort))
	}
	if c.Connections > 0 {
		cmd = append(cmd, "-c", strconv.Itoa(c.Connections))
	}
	if c.Threads > 0 {
		cmd = append(cmd, "-t", strconv.Itoa(c.Threads))
	}
	if c.LockPaged {
		cmd = append(cmd, "-k")
	}
	return cmd
}

func ParseConfigJsonData(data []byte) (*Config, error) {
	var cfg Config
	if err := json.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ParseConfigTomlData(data []byte) (*Config, error) {
	var cfg Config
	if _, err := toml.Decode(string(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	if strings.ToLower(filepath.Ext(fileName)) == ".toml" {
		return ParseConfigTomlData(data)
	} else {
		return ParseConfigJsonData(data)
	}
}
