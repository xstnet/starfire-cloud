package controllers

import (
	"fmt"
	"github.com/xstnet/starfire-cloud/internal/services/filemanager"
	"github.com/xstnet/starfire-cloud/internal/services/list"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

// 创建文件夹
func Mkdir(c *gin.Context) {
	userFile, err := filemanager.Mkdir(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, "创建成功", &gin.H{
		"id":         userFile.ID,
		"name":       userFile.Name,
		"is_dir":     1,
		"parent_id":  userFile.ParentId,
		"created_at": userFile.CreatedAt,
		"updated_at": userFile.UpdatedAt,
		"file": d.StringMap{
			"size": 0,
			"id":   0,
		},
	})
}

// 重命名
func Rename(c *gin.Context) {
	userFile, err := filemanager.Rename(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	// 需要返回名称，可能会有同名文件导致重命名，因此前端需要使用后端最新的数据
	response.Success(c, "重命名成功", &gin.H{"name": userFile.Name})
}

// Move 移动
// todo ：支持选择多个文件
func Move(c *gin.Context) {
	// 开启事物， 循环处理即可， 多文件移动的场景很少
	// 或先全部移过去， 再查询是否有重名的， 再根据ID去改名
	userFile, err := filemanager.Move(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c, "移动成功", &gin.H{"name": userFile.Name})
}

func List(c *gin.Context) {
	data, err := list.FileList(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OkWithData(c, data)
}

func Copy(c *gin.Context) {}

func Delete(c *gin.Context) {
	rowsAffected, err := filemanager.Delete(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c, fmt.Sprintf("删除成功, %d条数据已放入回收站", rowsAffected), nil)
}

func DirList(c *gin.Context) {
	data, err := list.DirList(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OkWithData(c, &gin.H{"list": data})
}
