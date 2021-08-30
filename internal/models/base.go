package models

import (
	"github.com/xstnet/starfire-cloud/internal/db"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `json:"id"`
	CreatedAt uint `json:"created_at"`
	UpdatedAt uint `json:"updated_at"`
}

func (model BaseModel) DB() *gorm.DB {
	return DB()
}

func DB() *gorm.DB {
	return db.DB
}
