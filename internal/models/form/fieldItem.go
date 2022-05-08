package form

// 基础字段

type ParentIdItem struct {
	ParentId uint `form:"parentId" json:"parentId" binding:"min=0"`
}

type FileIdItem struct {
	FileId uint `form:"fileId" json:"fileId" binding:"required,min=1"`
}

type NameItem struct {
	Name string `form:"name" json:"name" binding:"required,min=1"`
}

type Paging struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type SortItem struct {
	Sort uint8 `form:"sort"`
}
type OrderItem struct {
	Order string `form:"order"`
}

type FileIdsItem struct {
	FileIds []uint64 `json:"fileIds"`
}
