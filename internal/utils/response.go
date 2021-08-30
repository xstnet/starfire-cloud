package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/common"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Cost    string      `json:"cost"`
}

func ResponseOk(c *gin.Context, data interface{}) {
	ResponseJSON(c, common.CODE_SUCCESS, "ok", data)
}

func ResponseSuccess(c *gin.Context, message string, data interface{}) {
	ResponseJSON(c, common.CODE_SUCCESS, message, data)
}

func ResponseError(c *gin.Context, message string) {
	ResponseJSON(c, common.CODE_FAILURE, message, nil)
}

func ResponseJSON(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: message,
		Data:    data,
		Cost:    fmt.Sprintf("%v", (time.Since(c.GetTime("requestStartTime")))),
	})
}
