package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/services"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

func RecycleList(c *gin.Context)   {}
func RecycleDelete(c *gin.Context) {}
func RecycleRestore(c *gin.Context) {
	rowsAffected, err := services.RecycleRestore(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c, fmt.Sprintf("还原成功, %d条数据已恢复", rowsAffected), nil)
}
func RecycleClear(c *gin.Context) {}
