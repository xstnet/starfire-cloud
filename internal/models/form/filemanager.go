package form

// 获取文件列表
type FileList struct {
	Keyword  string `form:"keyword"`
	ParentId uint   `form:"parent_id"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Order    string `form:"order"`
	Sort     uint8  `form:"sort"`
}

type FileIdList struct {
	FIdList []uint64 `json:"fidlist"`
}

type FileCommon struct {
	FileIdList
}
