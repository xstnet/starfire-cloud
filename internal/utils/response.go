package utils

import (
	"github.com/xstnet/starfire-cloud/internal/common"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	// Cost    interface{} `json:"cost"`
}

func ResponseOk(data interface{}) *Response {
	return ResponseJSON(common.CODE_SUCCESS, "ok", data)
}

func ResponseSuccess(message string, data interface{}) *Response {
	return ResponseJSON(common.CODE_SUCCESS, message, data)
}

func ResponseError(message string) *Response {
	return ResponseJSON(common.CODE_FAILURE, message, nil)
}

func ResponseJSON(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
