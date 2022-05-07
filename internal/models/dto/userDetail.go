package dto

import (
	"github.com/xstnet/starfire-cloud/configs"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/pkg/systeminfo"
)

type UserDetail struct {
	Password string `json:"pass"`
	*models.User

	DiskInfo *systeminfo.DiskStatus `json:"DiskInfo"`
}

func LoadUserDetail(u *models.User) *UserDetail {
	// 处理总存储空间
	diskInfo := systeminfo.DiskInfo(configs.Upload.UploadRootPath)
	return &UserDetail{
		User:     u,
		DiskInfo: diskInfo,
	}
}
