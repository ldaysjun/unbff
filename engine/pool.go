package engine

import (
	"fmt"
	"github.com/ldaysjun/unbff/config"
	"github.com/ldaysjun/unbff/dao"
	"xorm.io/core"
	"xorm.io/xorm"
)

type nodePool struct {
	nodes map[int32]node
	db    *xorm.Engine
}

func NewNodePool() (*nodePool, error) {
	p := &nodePool{}
	p.db = newPoolDB(nil)
	nodes, err := dao.NodeList(p.db)
	if err != nil {
		return nil, err
	}
	for _, n := range nodes {
		sn, err := createNode(n)
		if err != nil {
			return nil, err
		}
		p.nodes[n.ID] = sn
	}
	return p, nil
}

func newPoolDB(conf *config.Config) *xorm.Engine {
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
