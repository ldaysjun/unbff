package model

import "time"

type SDLMetadata struct {
	ID     int32     `xorm:"c_id" json:"id"`         // pk auto incr
	App    string    `xorm:"c_app" json:"app"`       // Application
	Desc   string    `xorm:"c_desc" json:"desc"`     // schema definition language describe
	SDL    string    `xorm:"c_sdl" json:"sdl"`       // schema definition language
	Editor string    `xorm:"c_editor" json:"editor"` // editor
	Mtime  time.Time `xorm:"c_mtime" json:"mtime"`   // modify time
	Ctime  time.Time `xorm:"c_ctime" json:"ctime"`   // create time
}

// TableName table
func (s SDLMetadata) TableName() string {
	table := "t_dynamic_sdl_metadata"
	return table
}

type Node struct {
	ID       int32     `xorm:"c_id" json:"id"`                 // pk auto incr
	Source   string    `xorm:"c_source" json:"c_source"`       // data source
	Desc     string    `xorm:"c_desc" json:"c_desc"`           // describe
	NodeType int32     `xorm:"c_node_type" json:"c_node_type"` // node type
	Name     string    `xorm:"c_name" json:"c_name"`           // name
	Editor   string    `xorm:"c_editor" json:"c_editor"`       // editor
	Mtime    time.Time `xorm:"c_mtime" json:"mtime"`           // modify time
	Ctime    time.Time `xorm:"c_ctime" json:"ctime"`           // create time
}

// TableName table
func (n Node) TableName() string {
	table := "t_dynamic_source_node"
	return table
}
