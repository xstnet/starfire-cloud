package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/services/upload"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

// 秒传
func Instant(c *gin.Context) {
	data, err := upload.Instant(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OkWithData(c, &data)
}
