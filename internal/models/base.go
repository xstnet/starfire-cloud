package models

import (
	"github.com/xstnet/starfire-cloud/internal/db"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `json:"id"`
	CreatedAt uint `json:"created_at"`
	UpdatedAt uint `json:"updated_at"`
}

var scene = d.Map[string, []any]{
	"checkSameName": {"UserId", "ParentId", "IsDir", "Name"},     // 新建文件夹或上传文件时，检查同目录下是否有同名文件或文件夹
	"dirList":       {"UserId", "ParentId", "IsDir", "IsDelete"}, // 移动或复制文件时，获取当前级别的目录
}

// 是否删除
const (
	IsDeleteYes = 1
	IsDeleteNo  = 0
)

func DB() *gorm.DB {
	return db.DB
}

func (model *BaseModel) DB() *gorm.DB {
	return DB()
}

func (model *BaseModel) GetScene(key string) []any {
	return scene[key]
}
