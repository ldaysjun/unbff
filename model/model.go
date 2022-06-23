package model

type SDLMetadata struct {
	ID     int32  `xorm:"c_id" json:"id"`         // pk auto incr
	App    string `xorm:"c_app" json:"app"`       // Application
	Desc   string `xorm:"c_desc" json:"desc"`     // schema definition language describe
	SDL    string `xorm:"c_sdl" json:"sdl"`       // schema definition language
	Mtime  string `xorm:"c_mtime" json:"mtime"`   // modify time
	Ctime  string `xorm:"c_ctime" json:"ctime"`   // create time
	Editor string `xorm:"c_editor" json:"editor"` // editor
}

// TableName table
func (s SDLMetadata) TableName() string {
	table := "t_dynamic_sdl_metadata"
	return table
}

type Node struct {
	ID       int32  `xorm:"c_id" json:"id"`                 // pk auto incr
	Source   string `xorm:"c_source" json:"c_source"`       // data source
	Desc     string `xorm:"c_desc" json:"c_desc"`           // describe
	NodeType int32  `xorm:"c_node_type" json:"c_node_type"` // node type
	Name     string `xorm:"c_name" json:"c_name"`           // name
	Mtime    string `xorm:"c_mtime" json:"c_mtime"`         // modify time
	Ctime    string `xorm:"c_ctime" json:"c_ctime"`         // create time
	Editor   string `xorm:"c_editor" json:"c_editor"`       // editor
}

// TableName table
func (n Node) TableName() string {
	table := "t_dynamic_node"
	return table
}
