package proxy

import (
	"fmt"
	"github.com/juju/errors"
	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/router"
	"github.com/wandoulabs/cm/vt/tabletserver"
	"strings"
)

type Schema struct {
	db     string
	shards map[string]*Shard
	r      *router.Router
}

func (s *Server) parseSchemas() error {
	s.schemas = make(map[string]*Schema)

	for _, schemaCfg := range s.cfg.Schemas {
		db := strings.ToLower(schemaCfg.DB)
		if _, ok := s.schemas[db]; ok {
			return errors.Errorf("duplicate schema [%s].", schemaCfg.DB)
		}
		if len(schemaCfg.ShardIds) == 0 {
			return errors.Errorf("schema [%s] must have a shard.", schemaCfg.DB)
		}

		shards := make(map[string]*Shard)
		for _, n := range schemaCfg.ShardIds {
			if s.GetShard(n) == nil {
				return fmt.Errorf("schema [%s] shard [%s] config is not exists.", db, n)
			}

			if _, ok := shards[n]; ok {
				return fmt.Errorf("schema [%s] shard [%s] duplicate.", db, n)
			}
			shards[n] = s.GetShard(n)
		}

		r := router.NewRouter(&schemaCfg)
		schema := &Schema{
			db:     db,
			shards: shards,
			r:      r,
		}

		log.Infof("%+v", schema.r)
		s.schemas[db] = schema
	}

	return nil
}

func (s *Server) GetSchema(db string) *Schema {
	return s.schemas[db]
}

func (s *Server) MapToShards(db string, table string) []string {
	r := s.GetSchema(db).r
	shards := r.GetRule(table).MapToShards
	if len(shards) > 0 {
		return shards
	}

	return r.Default
}

func (s *Server) parseRowCacheCfg() tabletserver.RowCacheConfig {
	return s.cfg.RowCacheConf
}
