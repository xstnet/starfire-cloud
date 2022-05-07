package form

// FileList 获取文件列表
type FileList struct {
	Keyword string `form:"keyword"`
	ParentIdItem
	OrderItem
	SortItem
	Paging
}

type DirList struct {
	Paging
	ParentId uint
}

type DirTree struct {
	ParentId uint
}
