package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/services"
	"github.com/xstnet/starfire-cloud/internal/utils"
)

func RecycleList(c *gin.Context)   {}
func RecycleDelete(c *gin.Context) {}
func RecycleRestore(c *gin.Context) {
	rowsAffected, err := services.RecycleRestore(c, c.GetUint("userId"))
	if err != nil {
		utils.ResponseError(c, err.Error())
		return
	}
	utils.ResponseSuccess(c, fmt.Sprintf("还原成功, %d条数据已恢复", rowsAffected), nil)
}
func RecycleClear(c *gin.Context) {}
