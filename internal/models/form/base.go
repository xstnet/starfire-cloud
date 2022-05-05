package form

import "github.com/gin-gonic/gin"

type Base struct {
}

func GetForm[T any](c *gin.Context) (*T, error) {
	form := new(T)
	err := c.ShouldBind(form)
	return form, err
}

func GetJsonForm[T any](c *gin.Context) (*T, error) {
	jsonForm := new(T)
	err := c.ShouldBindJSON(jsonForm)
	return jsonForm, err
}
