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

// SDLMetadataList sdl metadata list
func SDLMetadataList(db *xorm.Engine) ([]*model.SDLMetadata, error) {
	var sdlList []*model.SDLMetadata
	if err := db.Find(&sdlList); err != nil {
		return nil, err
	}
	return sdlList, nil
}
