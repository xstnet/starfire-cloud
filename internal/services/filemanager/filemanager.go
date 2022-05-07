package filemanager

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/pkg/helper/convert"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
	"github.com/xstnet/starfire-cloud/pkg/helper/dirHelper"
	"github.com/xstnet/starfire-cloud/pkg/helper/fileHelper"
)

// Mkdir 创建文件夹
func Mkdir(c *gin.Context, userId uint) (*models.UserFile, error) {
	params, err := form.GetJsonForm[form.Mkdir](c)
	if err != nil {
		return nil, err
	}

	if err := dirHelper.CheckName(params.Name); err != nil {
		return nil, err
	}

	userFile := &models.UserFile{
		UserId:   userId,
		ParentId: params.ParentId,
		IsDir:    models.IsDirYes,
		Name:     params.Name,
	}

	if err := userFile.Mkdir(); err != nil {
		return nil, errors.New("创建文件夹失败，原因: " + err.Error())
	}

	return userFile, nil
}

// Rename 重命名
func Rename(c *gin.Context, userId uint) (*models.UserFile, error) {
	params, err := form.GetJsonForm[form.Rename](c)
	if err != nil {
		return nil, errors.InvalidParameter()
	}

	userFileModel := new(models.UserFile)
	if result := userFileModel.DB().First(userFileModel, params.FileId); result.Error != nil || userFileModel.UserId != userId {
		return nil, errors.New("操作对象不存在")
	}

	if userFileModel.IsDir == models.IsDirNo {
		if err := fileHelper.CheckName(params.NewName); err != nil {
			return nil, err
		}
	} else {
		if err := dirHelper.CheckName(params.NewName); err != nil {
			return nil, err
		}
	}

	if err := userFileModel.Rename(params.NewName); err != nil {
		return nil, err
	}

	return userFileModel, nil
}

// Move 移动
func Move(c *gin.Context, userId uint) (*models.UserFile, error) {
	var data = make(d.StringMap, 4)
	err := c.ShouldBindJSON(&data)

	if err != nil {
		return nil, errors.InvalidParameter()
	}

	fromId, ok := convert.GetFloat64(data["from_id"])
	if !ok || fromId <= 0 {
		return nil, errors.InvalidParameter()
	}
	destId, ok := convert.GetFloat64(data["dest_id"])
	if !ok || destId <= 0 {
		return nil, errors.InvalidParameter()
	}

	userFile := &models.UserFile{}
	if result := userFile.DB().First(userFile, uint(fromId)); result.Error != nil || userFile.UserId != userId {
		return nil, errors.New("操作对象不存在")
	}

	userFile.ParentId = uint(destId)
	if err := userFile.Move(); err != nil {
		return nil, err
	}

	return userFile, nil
}

// Delete 删除
// 只标记当前节点，不处理子元素，从回收站删除时再处理子元素
func Delete(c *gin.Context, userId uint) (int64, error) {
	params, err := form.GetJsonForm[form.FileIdsItem](c)
	if err != nil {
		return 0, errors.InvalidParameter()
	}

	result := models.DB().Model(new(models.UserFile)).
		Where(params.FileIds).
		Where("user_id = ? and is_delete=?", userId, models.IsDeleteNo).
		Update("is_delete", 1)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
