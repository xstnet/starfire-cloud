package models

import (
	"time"

	"github.com/xstnet/starfire-cloud/internal/errors"
)

type UserFile struct {
	BaseModel
	UserId   uint   `json:"user_id"`
	ParentId uint   `json:"parent_id"`
	FileId   uint   `json:"file_id"`
	IsDir    uint8  `json:"is_dir"`
	IsDelete uint8  `json:"is_delete"`
	Name     string `json:"name"`
}

// 是否文件夹
const (
	IS_DIR_YES = 1
	IS_DIR_NO  = 0
)

const (
	IS_DELETE_YES = 1
	IS_DELETE_NO  = 0
)

var scene = map[string][]interface{}{
	"check_samename": {"UserId", "ParentId", "IsDir", "Name"}, // 新建文件夹或上传文件时，检查同目录下是否有同名文件或文件夹
}

func (uf *UserFile) GetScene(key string) []interface{} {
	return scene[key]
}

// 创建文件夹
func (uf *UserFile) Mkdir() error {
	if err := uf.checkParentId(); err != nil {
		return err
	}
	// 防止 shoudBindJson从外部传参，初始化
	{
		uf.IsDir = IS_DIR_YES
		uf.IsDelete = IS_DELETE_NO
		uf.CreatedAt, uf.UpdatedAt, uf.ID = 0, 0, 0
	}

	uf.processSameName()

	result := uf.DB().Create(uf)
	return result.Error
}

func (uf *UserFile) Rename() error {
	uf.processSameName()
	result := uf.DB().Model(uf).Update("Name", uf.Name)
	return result.Error
}

// 移动
func (uf *UserFile) Move() error {
	if err := uf.checkParentId(); err != nil {
		return err
	}

	uf.processSameName()
	result := uf.DB().Model(uf).Select("parent_id", "name").Updates(uf)
	return result.Error
}

// 创建文件夹前先校验归属
func (uf *UserFile) checkParentId() error {
	// 不是在根目录创建，需要验证归属文件夹是否属于当前用户
	if uf.ParentId > 0 {
		var count int64
		uf.DB().Model(uf).Where("id = ? AND user_id = ? AND is_dir = ?", uf.ParentId, uf.UserId, IS_DIR_YES).Count(&count)
		if count == 0 {
			return errors.New("归属文件夹不存在")
		}
	}

	return nil
}

// 处理同名文件或文件夹，使用日期时间后缀来避免
// 如果所选目录下已有同名的文件，加上日期时间做为后缀重命名
// 不考虑在回收站的文件，一样当做同名处理
func (uf *UserFile) processSameName() {
	var count int64
	uf.DB().Model(uf).Where(uf, uf.GetScene("check_samename")...).Count(&count)
	if count > 0 {
		// eg: file_20210830_234812
		// 若是同一秒内处理多个请求，可能会导致新文件名依然重复，由于是面象私有云，为性能考虑不处理
		uf.Name += "_" + time.Unix(int64(time.Now().Unix()), 0).Format("20060102_150405")
	}
}
