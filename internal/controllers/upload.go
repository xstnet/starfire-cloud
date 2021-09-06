package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/services"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

// 暂不使用
func BatchUpload(c *gin.Context) {
	err := services.BatchUpload(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Ok(c)
}

func UploadFile(c *gin.Context) {
	data, err := services.UploadFile(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, "上传失败, "+err.Error())
		return
	}
	response.Success(c, "上传成功", &data)
}

func PreUpload(c *gin.Context) {
	data, err := services.PreUpload(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OkWithData(c, &data)
}

// 秒传
func Instant(c *gin.Context) {
	data, err := services.PreUpload(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OkWithData(c, &data)
}
