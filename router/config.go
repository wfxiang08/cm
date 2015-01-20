package router

import "github.com/wandoulabs/cm/config"

var (
	DefaultRuleType = "default"
	HashRuleType    = "hash"
	RangeRuleType   = "range"
)

type RuleConfig struct {
	config.ShardConfig
}

func (c *RuleConfig) ParseRule(db string) (*Rule, error) {
	r := &Rule{
		DB:          db,
		Table:       c.Table,
		ShardingKey: c.ShardingKey,
		Shard:       c.Shard,
	}

	if err := c.parseShard(r); err != nil {
		return nil, err
	}

	return r, nil
}

/*
func (c *RuleConfig) parseShards(r *Rule) error {
	// Note: did not used yet, by HuangChuanTong
	reg, err := regexp.Compile(`(\w+)\((\d+)\-(\d+)\)`)
	if err != nil {
		return err
	}

	ns := c.Shard

	shards := map[string]struct{}{}

	for _, n := range ns {
		n = strings.TrimSpace(n)
		if s := reg.FindStringSubmatch(n); s == nil {
			if _, ok := shards[n]; ok {
				return fmt.Errorf("duplicate node %s", n)
			}

			shards[n] = struct{}{}
			r.Shard = append(r.Shards, n)
		} else {
			var start, stop int
			if start, err = strconv.Atoi(s[2]); err != nil {
				return err
			}

			if stop, err = strconv.Atoi(s[3]); err != nil {
				return err
			}

			if start >= stop {
				return fmt.Errorf("invalid node format %s", n)
			}

			for i := start; i <= stop; i++ {
				n = fmt.Sprintf("%s%d", s[1], i)

				if _, ok := shards[n]; ok {
					return fmt.Errorf("duplicate node %s", n)
				}

				shards[n] = struct{}{}
				r.Shards = append(r.Shards, n)

			}
		}
	}

	if len(r.Shards) == 0 {
		return fmt.Errorf("empty shards info")
	}

	if r.Type == DefaultRuleType && len(r.Shards) != 1 {
		return fmt.Errorf("default rule must have only one node")
	}

	return nil
}
*/

func (c *RuleConfig) parseShard(r *Rule) error {
	return nil
}
