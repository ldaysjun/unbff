package engine

import "xorm.io/xorm"

type dbNode struct {
	xdb  *xorm.Engine
	name string
}

func (n *dbNode) do() (interface{}, error) {
	return nil, nil
}
