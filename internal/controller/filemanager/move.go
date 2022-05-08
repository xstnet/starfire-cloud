package filemanager

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/pkg/helper/convert"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

// Move 移动
// todo ：支持选择多个文件
func Move(c *gin.Context) {
	// 开启事物， 循环处理即可， 多文件移动的场景很少
	// 或先全部移过去， 再查询是否有重名的， 再根据ID去改名
	userFile, err := doMove(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c, "移动成功", &d.StringMap{"name": userFile.Name})
}

// Move 移动
func doMove(c *gin.Context, userId uint) (*models.UserFile, error) {
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
