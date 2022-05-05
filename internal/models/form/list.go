package form

type Paging struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

// FileList 获取文件列表
type FileList struct {
	Keyword  string `form:"keyword"`
	ParentId uint   `form:"parentId"`
	Order    string `form:"order"`
	Sort     uint8  `form:"sort"`
	Paging
}

type DirList struct {
	Paging
	ParentId uint
}
