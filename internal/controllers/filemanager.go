package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/services"
	"github.com/xstnet/starfire-cloud/internal/utils"
)

// 创建文件夹
func Mkdir(c *gin.Context) {
	userFile, err := services.Mkdir(c, c.GetUint("userId"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	utils.ResponseSuccess(c, "创建成功", &gin.H{
		"id":   userFile.ID,
		"name": userFile.Name,
	})
}

// 重命名
func Rename(c *gin.Context) {
	userFile, err := services.Rename(c, c.GetUint("userId"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}
	// 需要返回名称，可能会有同名文件导致重命名，因此前端需要使用后端最新的数据
	utils.ResponseSuccess(c, "重命名成功", &gin.H{"name": userFile.Name})
}

// 移动
func Move(c *gin.Context) {
	userFile, err := services.Move(c, c.GetUint("userId"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}
	utils.ResponseSuccess(c, "移动成功", &gin.H{"name": userFile.Name})
}
