package dao

import (
	"github.com/ldaysjun/unbff/model"
	"xorm.io/xorm"
)

// NodeList get a list of resource nodes
func NodeList(db *xorm.Engine) ([]*model.Node, error) {
	var nodes []*model.Node
	if err := db.Find(&nodes); err != nil {
		return nil, err
	}
	return nodes, nil
}

func SDLMetadataList(db *xorm.Engine) ([]*model.SDLMetadata, error) {
	var sdls []*model.SDLMetadata
	if err := db.Find(&sdls); err != nil {
		return nil, err
	}
	return sdls, nil
}
