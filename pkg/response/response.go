package response

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CODE_SUCCESS = 0     // 成功
	CODE_FAILURE = 1     // 失败
	CODE_RELOGIN = 27149 // 重新登录
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Cost    string      `json:"cost"`
}

func Ok(c *gin.Context) {
	JSON(c, CODE_SUCCESS, "ok", nil)
}

func OkWithData(c *gin.Context, data interface{}) {
	JSON(c, CODE_SUCCESS, "ok", data)
}

func Success(c *gin.Context, message string, data interface{}) {
	JSON(c, CODE_SUCCESS, message, data)
}

func Error(c *gin.Context, message string) {
	JSON(c, CODE_FAILURE, message, nil)
}

func ErrorWithData(c *gin.Context, message string, data interface{}) {
	JSON(c, CODE_FAILURE, message, data)
}

func JSON(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: message,
		Data:    data,
		Cost:    fmt.Sprintf("%v", (time.Since(c.GetTime("requestStartTime")))),
	})
}
