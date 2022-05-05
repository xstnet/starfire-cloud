package list

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
)

// DirList 文件夹列表
func DirList(c *gin.Context, userId uint) (*[]models.UserFile, error) {
	listForm, err := form.GetForm[form.DirList](c)
	if err != nil {
		return nil, errors.InvalidParameter()
	}

	userFile := &models.UserFile{
		UserId:   userId,
		ParentId: listForm.ParentId,
	}

	data := userFile.DirList()

	return data, nil
}
