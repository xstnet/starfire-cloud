package models

type UserFile struct {
	BaseField
	ParentId uint   `json:"parent_id"`
	UserId   uint   `json:"user_id"`
	FileId   uint   `json:"file_id"`
	IsDir    uint8  `json:"is_dir"`
	IsDelete uint8  `json:"is_delete"`
	Name     string `json:"name"`
}
