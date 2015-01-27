package router

import (
	"github.com/wandoulabs/cm/config"
)

type Router struct {
	All     map[string]*config.TableRule //table name -> rules
	Default []string
}

func NewRouter(cfg *config.SchemaConfig) *Router {
	r := &Router{All: make(map[string]*config.TableRule)}
	for _, tr := range cfg.RouterConifg.TableRule {
		tmp := &config.TableRule{}
		*tmp = tr
		r.All[tr.Table] = tmp
	}

	r.Default = cfg.RouterConifg.Default

	return r
}

func (r *Router) GetRule(table string) *config.TableRule {
	return r.All[table]
}
