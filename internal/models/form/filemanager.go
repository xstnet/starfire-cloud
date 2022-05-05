package form

type FileIdList struct {
	FIdList []uint64 `json:"fidlist"`
}

type FileCommon struct {
	FileIdList
}
