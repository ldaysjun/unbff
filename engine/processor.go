package engine

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/ldaysjun/unbff/config"
	"xorm.io/core"
	"xorm.io/xorm"
)

type processor struct {
	db *xorm.Engine
}

func newProcessor(conf *config.Config) *processor {
	p := &processor{}
	p.db = newProcessorDB(conf)
	return nil
}

func (p *processor) generateFields() (map[string]graphql.Fields, error) {
	
	return nil, nil
}

func newProcessorDB(conf *config.Config) *xorm.Engine {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		conf.PoolDB.DBUser, conf.PoolDB.DBPwd, conf.PoolDB.DBHost, conf.PoolDB.DBPort,
		conf.PoolDB.DBName, conf.PoolDB.DBCharset)
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	engine.SetMapper(core.SameMapper{})
	if err := engine.Ping(); err != nil {
		panic(err)
	}
	return engine
}
