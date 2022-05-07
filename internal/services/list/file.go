package list

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
)

// FileList 获取文件列表
func FileList(c *gin.Context, userId uint) (*d.StringMap, error) {
	listForm, err := form.GetForm[form.FileList](c)
	if err != nil {
		return nil, err
	}
	return &d.StringMap{"a": listForm}, nil

}
