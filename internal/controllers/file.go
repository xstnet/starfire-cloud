package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/services"
	"github.com/xstnet/starfire-cloud/internal/utils"
)

func Mkdir(c *gin.Context) {
	data := services.Mkdir(c)
	utils.ResponseOk(c, c.GetInt("userId"))
}
