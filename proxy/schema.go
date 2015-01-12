package proxy

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/wandoulabs/cm/router"
	"github.com/wandoulabs/cm/vt/tabletserver"
	"strings"
)

type Schema struct {
	db    string
	nodes map[string]*Node
	rule  *router.Router
}

func (s *Server) parseSchemas() error {
	s.schemas = make(map[string]*Schema)

	for _, schemaCfg := range s.cfg.Schemas {
		db := strings.ToLower(schemaCfg.DB)
		if _, ok := s.schemas[db]; ok {
			return errors.Errorf("duplicate schema [%s].", schemaCfg.DB)
		}
		if len(schemaCfg.Nodes) == 0 {
			return errors.Errorf("schema [%s] must have a node.", schemaCfg.DB)
		}

		nodes := make(map[string]*Node)
		for _, n := range schemaCfg.Nodes {
			if s.GetNode(n) == nil {
				return fmt.Errorf("schema [%s] node [%s] config is not exists.", db, n)
			}

			if _, ok := nodes[n]; ok {
				return fmt.Errorf("schema [%s] node [%s] duplicate.", db, n)
			}
			nodes[n] = s.GetNode(n)
		}

		rule, err := router.NewRouter(&schemaCfg)
		if err != nil {
			return err
		}

		s.schemas[db] = &Schema{
			db:    db,
			nodes: nodes,
			rule:  rule,
		}
	}

	return nil
}

func (s *Server) GetSchema(db string) *Schema {
	return s.schemas[db]
}

func (s *Server) parseRowCacheCfg() tabletserver.RowCacheConfig {
	return s.cfg.RowCacheConf
}
