package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/services/upload"
	"github.com/xstnet/starfire-cloud/pkg/response"
	"strconv"
)

func SingleUpload(c *gin.Context) {
	targetId, err := strconv.Atoi(c.PostForm("target_id"))
	// 校验参数
	if err != nil || targetId < 0 {
		response.Error(c, "上传失败, "+err.Error())
		return
	}
	data, err := upload.SingleUpload(c, c.GetUint("userId"), targetId)
	if err != nil {
		response.Error(c, "上传失败, "+err.Error())
		return
	}
	response.Success(c, "上传成功", &data)
}
