package router

import (
	"fmt"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/config"
)

type Rule struct {
	DB          string
	Table       string
	ShardingKey string
	Type        string
	Shard       string
}

func (r *Rule) GetShard(key interface{}) string {
	return r.Shard
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s.%s?shardingkey=%v&shard=%s",
		r.DB, r.Table, r.ShardingKey, r.Shard)
}

func NewDefaultRule(db string, shard string) *Rule {
	var r *Rule = &Rule{
		DB:    db,
		Shard: shard,
	}

	log.Warningf("%+v", r)
	return r
}

func (r *Router) GetRule(table string) *Rule {
	rule := r.Rules[table]
	if rule == nil {
		return r.DefaultRule
	} else {
		if len(rule.Shard) == 0 {
			return r.DefaultRule
		}

		return rule
	}
}

type Router struct {
	DB          string
	Rules       map[string]*Rule //key is <table name>
	DefaultRule *Rule
	shards      []string //just for human saw
}

func NewRouter(schemaConfig *config.SchemaConfig) (*Router, error) {
	if !includeNode(schemaConfig.Shards, schemaConfig.RulesConifg.Default) {
		return nil, fmt.Errorf("default shard[%s] not in the shards list.",
			schemaConfig.RulesConifg.Default)
	}

	rt := &Router{
		DB:     schemaConfig.DB,
		shards: schemaConfig.Shards,
		Rules:  make(map[string]*Rule, len(schemaConfig.RulesConifg.ShardRule)),
	}

	rt.DefaultRule = NewDefaultRule(rt.DB, schemaConfig.RulesConifg.Default)

	for _, shard := range schemaConfig.RulesConifg.ShardRule {
		rc := &RuleConfig{shard}
		rule, err := rc.ParseRule(rt.DB)
		if err != nil {
			return nil, err
		}

		if rule.Type == DefaultRuleType {
			return nil, fmt.Errorf("[default-rule] duplicate, must only one.")
		} else {
			if _, ok := rt.Rules[rule.Table]; ok {
				return nil, fmt.Errorf("table %s rule in %s duplicate", rule.Table, rule.DB)
			}
			rt.Rules[rule.Table] = rule
		}
	}

	log.Warningf("%+v", rt.DefaultRule)
	return rt, nil
}

func includeNode(shards []string, shard string) bool {
	for _, n := range shards {
		if n == shard {
			return true
		}
	}

	return false
}
