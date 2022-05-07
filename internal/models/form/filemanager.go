package form

// 基础字段

type PatientIdItem struct {
	PatientId uint `binding:"required,min=0"`
}

type NameItem struct {
	Name string `binding:"required,min=1"`
}

type FileIdList struct {
	FIdList []uint64 `json:"fidlist"`
}

type FileCommon struct {
	FileIdList
}

type Mkdir struct {
	NameItem
	PatientIdItem
}
