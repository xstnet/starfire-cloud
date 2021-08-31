package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/common"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/utils"
)

// 创建文件夹
func Mkdir(c *gin.Context, userId uint) (*models.UserFile, error) {
	userFile := &models.UserFile{}
	err := c.ShouldBindJSON(&userFile)
	if err != nil {
		return nil, errors.InvalidParameter()
	}

	if err := common.CheckDirname(userFile.Name); err != nil {
		return nil, err
	}

	userFile.UserId = userId

	if err := userFile.Mkdir(); err != nil {
		return nil, errors.New("创建文件夹失败，原因: " + err.Error())
	}

	return userFile, nil
}

// 重命名
func Rename(c *gin.Context, userId uint) (*models.UserFile, error) {
	var data = make(gin.H, 4)
	err := c.ShouldBindJSON(&data)

	if err != nil {
		return nil, errors.InvalidParameter()
	}

	id, ok := utils.GetFloat64(data["id"])
	if !ok || id <= 0 {
		return nil, errors.InvalidParameter()
	}
	newname, ok := utils.GetString(data["newname"])
	if !ok {
		return nil, errors.InvalidParameter()
	}

	if err := common.CheckFilename(newname); err != nil {
		return nil, err
	}

	userFile := &models.UserFile{}
	if result := userFile.DB().First(userFile, uint(id)); result.Error != nil || userFile.UserId != userId {
		return nil, errors.New("操作对象不存在")
	}

	userFile.Name = newname
	if err := userFile.Rename(); err != nil {
		return nil, err
	}

	return userFile, nil
}

// 移动
func Move(c *gin.Context, userId uint) (*models.UserFile, error) {
	var data = make(gin.H, 4)
	err := c.ShouldBindJSON(&data)

	if err != nil {
		return nil, errors.InvalidParameter()
	}

	fromId, ok := utils.GetFloat64(data["from_id"])
	if !ok || fromId <= 0 {
		return nil, errors.InvalidParameter()
	}
	destId, ok := utils.GetFloat64(data["dest_id"])
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
