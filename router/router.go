package router

import (
	"fmt"

	"github.com/wandoulabs/cm/config"
)

type Rule struct {
	DB    string
	Table string
	Key   string

	Type string

	Node string
}

func (r *Rule) GetNode(key interface{}) string {
	return r.Node
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s.%s?key=%v&node=%s",
		r.DB, r.Table, r.Key, r.Node)
}

func NewDefaultRule(db string, node string) *Rule {
	var r *Rule = &Rule{
		DB: db,
	}
	return r
}

func (r *Router) GetRule(table string) *Rule {
	rule := r.Rules[table]
	if rule == nil {
		return r.DefaultRule
	} else {
		return rule
	}
}

type Router struct {
	DB          string
	Rules       map[string]*Rule //key is <table name>
	DefaultRule *Rule
	nodes       []string //just for human saw
}

func NewRouter(schemaConfig *config.SchemaConfig) (*Router, error) {
	if !includeNode(schemaConfig.Nodes, schemaConfig.RulesConifg.Default) {
		return nil, fmt.Errorf("default node[%s] not in the nodes list.",
			schemaConfig.RulesConifg.Default)
	}

	rt := new(Router)
	rt.DB = schemaConfig.DB
	rt.nodes = schemaConfig.Nodes
	rt.Rules = make(map[string]*Rule, len(schemaConfig.RulesConifg.ShardRule))
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
	return rt, nil
}

func includeNode(nodes []string, node string) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}

	return false
}
