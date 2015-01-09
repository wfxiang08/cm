package proxy

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/wandoulabs/cm/router"
	"github.com/wandoulabs/cm/vt/tabletserver"
)

type Schema struct {
	db    string
	nodes map[string]*Node
	rule  *router.Router
}

func (s *Server) parseSchemas() error {
	s.schemas = make(map[string]*Schema)

	for _, schemaCfg := range s.cfg.Schemas {
		if _, ok := s.schemas[schemaCfg.DB]; ok {
			return errors.Errorf("duplicate schema [%s].", schemaCfg.DB)
		}
		if len(schemaCfg.Nodes) == 0 {
			return errors.Errorf("schema [%s] must have a node.", schemaCfg.DB)
		}

		nodes := make(map[string]*Node)
		for _, n := range schemaCfg.Nodes {
			if s.GetNode(n) == nil {
				return fmt.Errorf("schema [%s] node [%s] config is not exists.", schemaCfg.DB, n)
			}

			if _, ok := nodes[n]; ok {
				return fmt.Errorf("schema [%s] node [%s] duplicate.", schemaCfg.DB, n)
			}
			nodes[n] = s.GetNode(n)
		}

		rule, err := router.NewRouter(&schemaCfg)
		if err != nil {
			return err
		}

		s.schemas[schemaCfg.DB] = &Schema{
			db:    schemaCfg.DB,
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
