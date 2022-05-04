package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/services"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

func SingleUpload(c *gin.Context) {
	data, err := services.UploadFile(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, "上传失败, "+err.Error())
		return
	}
	response.Success(c, "上传成功", &data)
}
