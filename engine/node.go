package engine

import (
	"errors"
	"github.com/ldaysjun/unbff/model"
	"xorm.io/core"
	"xorm.io/xorm"
)

const (
	DefaultNodeType int = 0
	DBNodeType          = 100
)

type node interface {
	do() (interface{}, error)
}

// createNode 创建节点
func createNode(m *model.Node) (node, error) {
	switch m.NodeType {
	case DBNodeType:
		db, err := xorm.NewEngine("mysql", m.Source)
		if err != nil {
			return nil, err
		}
		db.SetMapper(core.SameMapper{})
		db.ShowSQL(true)
		if err := db.Ping(); err != nil {
			return nil, err
		}
		return &dbNode{xdb: db, name: m.Name}, nil
	default:
		return nil, errors.New("")
	}
}
