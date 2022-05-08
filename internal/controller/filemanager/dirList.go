package filemanager

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

func DirList(c *gin.Context) {
	data, err := doList(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.OkWithData(c, data)
}

// FileList 获取文件列表
func doDirList(c *gin.Context, userId uint) (*d.StringMap, error) {
	listForm, err := form.GetForm[form.FileList](c)
	if err != nil {
		return nil, err
	}
	return &d.StringMap{"a": listForm}, nil

}
