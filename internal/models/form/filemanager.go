package form

// Mkdir 创建文件夹
type Mkdir struct {
	NameItem
	ParentIdItem
}

// Rename 重命名
type Rename struct {
	NewName string `binding:"required,min=1"`
	FileIdItem
}
