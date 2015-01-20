package proxy

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/wandoulabs/cm/router"
	"github.com/wandoulabs/cm/vt/tabletserver"
	"strings"
)

type Schema struct {
	db     string
	shards map[string]*Shard
	rule   *router.Router
}

func (s *Server) parseSchemas() error {
	s.schemas = make(map[string]*Schema)

	for _, schemaCfg := range s.cfg.Schemas {
		db := strings.ToLower(schemaCfg.DB)
		if _, ok := s.schemas[db]; ok {
			return errors.Errorf("duplicate schema [%s].", schemaCfg.DB)
		}
		if len(schemaCfg.Shards) == 0 {
			return errors.Errorf("schema [%s] must have a shard.", schemaCfg.DB)
		}

		shards := make(map[string]*Shard)
		for _, n := range schemaCfg.Shards {
			if s.GetShard(n) == nil {
				return fmt.Errorf("schema [%s] shard [%s] config is not exists.", db, n)
			}

			if _, ok := shards[n]; ok {
				return fmt.Errorf("schema [%s] shard [%s] duplicate.", db, n)
			}
			shards[n] = s.GetShard(n)
		}

		rule, err := router.NewRouter(&schemaCfg)
		if err != nil {
			return err
		}

		s.schemas[db] = &Schema{
			db:     db,
			shards: shards,
			rule:   rule,
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
