package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/wandoulabs/cm/vt/tabletserver"
)

type SchemaConfig struct {
	DB           string       `json:"db"`
	ShardIds     []string     `json:"shard_ids"`
	RouterConifg RouterConfig `json:"router"`
	CacheSize    int          `json:"cache_size,string"` //m
}

type RouterConfig struct {
	Default   string      `json:"default"`
	TableRule []TableRule `json:"table_rules"`
}

type TableRule struct {
	Table        string `json:"table"`
	ShardingKey  string `json:"key"`
	RowCacheType string `json:"row_cache_type"`
	MapToShard   string `json:"map_to_shards"` //shard ids
}

type ShardConfig struct {
	Id       string `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`

	Master string `json:"master"`
	Slave  string `json:"slave"`
}

type Config struct {
	Addr         string                      `json:"addr"`
	User         string                      `json:"user"`
	Password     string                      `json:"password"`
	LogLevel     string                      `json:"log_level"`
	SkipAuth     bool                        `json:"skip_auth"`
	Shards       []ShardConfig               `json:"shards"`
	Schemas      []SchemaConfig              `json:"schemas"`
	RowCacheConf tabletserver.RowCacheConfig `json:"rowcache_conf"`
}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := json.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(data)
}
