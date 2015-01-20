package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/wandoulabs/cm/vt/tabletserver"
)

type SchemaConfig struct {
	DB          string      `json:"db"`
	Shards      []string    `json:"shards"`
	RulesConifg RulesConfig `json:"rules"`
	CacheSize   int         `json:"cache_size,string"` //m
}

type RulesConfig struct {
	Default   string        `json:"default"`
	ShardRule []ShardConfig `json:"shard"`
}

type ShardConfig struct {
	Table        string `json:"table"`
	ShardingKey  string `json:"key"`
	RowCacheType string `json:"row_cache_type"`
	Shard        string `json:"shard"`

	Name    string `json:"name"`
	RWSplit bool   `json:"rw_split"`

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
