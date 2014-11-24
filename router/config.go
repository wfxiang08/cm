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
	r := new(Rule)
	r.DB = db
	r.Table = c.Table
	r.Key = c.Key
	r.Node = c.Node

	if err := c.parseShard(r); err != nil {
		return nil, err
	}

	return r, nil
}

/*
func (c *RuleConfig) parseNodes(r *Rule) error {
	// Note: did not used yet, by HuangChuanTong
	reg, err := regexp.Compile(`(\w+)\((\d+)\-(\d+)\)`)
	if err != nil {
		return err
	}

	ns := c.Node

	nodes := map[string]struct{}{}

	for _, n := range ns {
		n = strings.TrimSpace(n)
		if s := reg.FindStringSubmatch(n); s == nil {
			if _, ok := nodes[n]; ok {
				return fmt.Errorf("duplicate node %s", n)
			}

			nodes[n] = struct{}{}
			r.Node = append(r.Nodes, n)
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

				if _, ok := nodes[n]; ok {
					return fmt.Errorf("duplicate node %s", n)
				}

				nodes[n] = struct{}{}
				r.Nodes = append(r.Nodes, n)

			}
		}
	}

	if len(r.Nodes) == 0 {
		return fmt.Errorf("empty nodes info")
	}

	if r.Type == DefaultRuleType && len(r.Nodes) != 1 {
		return fmt.Errorf("default rule must have only one node")
	}

	return nil
}
*/

func (c *RuleConfig) parseShard(r *Rule) error {
	return nil
}
