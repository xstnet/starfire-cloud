package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/services/upload"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

// BatchUpload 暂不使用
func BatchUpload(c *gin.Context) {
	err := upload.BatchUpload(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Ok(c)
}
