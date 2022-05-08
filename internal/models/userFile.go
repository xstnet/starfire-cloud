package models

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/xstnet/starfire-cloud/internal/errors"
)

type UserFile struct {
	BaseModel
	UserId   uint   `json:"userId"`
	ParentId uint   `json:"parentId"`
	FileId   uint   `json:"fileId"`
	IsDir    uint8  `json:"isDir"`
	IsDelete uint8  `json:"isDelete"`
	Name     string `json:"name"`
}

// 是否文件夹
const (
	IsDirYes = 1
	IsDirNo  = 0
)

//func (uf *UserFile) GetScene(key string) []string {
//	return scene[key]
//}

// Mkdir 创建文件夹
func (uf *UserFile) Mkdir() error {
	if err := uf.CheckParentId(); err != nil {
		return err
	}
	uf.processSameName()

	result := uf.DB().Create(uf)
	return result.Error
}

func (uf *UserFile) Rename(newName string) error {
	uf.Name = newName
	uf.processSameName()
	result := uf.DB().Model(uf).Update("Name", uf.Name)
	return result.Error
}

// Move 移动
func (uf *UserFile) Move() error {
	if err := uf.CheckParentId(); err != nil {
		return err
	}

	uf.processSameName()
	result := uf.DB().Model(uf).Select("parent_id", "name").Updates(uf)
	return result.Error
}

// 创建文件夹前先校验归属
func (uf *UserFile) CheckParentId() error {
	// 不是在根目录创建，需要验证归属文件夹是否属于当前用户
	if uf.ParentId > 0 {
		var count int64
		uf.DB().Model(uf).Where("id = ? AND user_id = ? AND is_dir = ?", uf.ParentId, uf.UserId, IsDirYes).Count(&count)
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
	uf.DB().Model(uf).Where(uf, uf.GetScene("checkSameName")...).Count(&count)
	if count <= 0 {
		return
	}

	// 如果是文件， 需要处理一下扩展名
	var ext string
	if uf.IsDir == IsDirNo {
		ext = filepath.Ext(uf.Name)
		uf.Name = strings.TrimSuffix(uf.Name, ext)
	}
	// eg: file_20210830_234812
	// 若是同一秒内处理多个请求，可能会导致新文件名依然重复，由于是面象私有云，为性能考虑不处理
	uf.Name += "_" + time.Unix(int64(time.Now().Unix()), 0).Format("20060102_150405")
	// 针对文件添加扩展名
	uf.Name += ext
}

func (uf *UserFile) BindFile() error {
	uf.processSameName()

	return uf.DB().Create(uf).Error
}

func (uf *UserFile) DirList() *[]UserFile {
	uf.IsDelete = IsDeleteNo
	uf.IsDir = IsDirYes

	// 预定义10个容量
	var data = make([]UserFile, 0, 10)
	uf.DB().Model(uf).Where(uf, uf.GetScene("dirList")...).Order("updated_at desc").Find(&data)

	return &data
}
