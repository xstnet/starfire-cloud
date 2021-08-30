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
		"id": userFile.ID,
	})
}
