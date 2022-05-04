package list

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
)

func DirList(c *gin.Context, userId uint) (*[]models.UserFile, error) {
	listForm := &form.FileList{}
	if err := c.ShouldBind(listForm); err != nil {
		return nil, errors.InvalidParameter()
	}

	userFile := &models.UserFile{
		UserId:   userId,
		ParentId: listForm.ParentId,
	}

	data := userFile.DirList()

	return data, nil
}
