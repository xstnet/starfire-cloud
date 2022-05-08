package filemanager

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/dto"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/pkg/helper/dirHelper"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

// Mkdir 创建文件夹
func Mkdir(c *gin.Context) {
	userId := c.GetUint("userId")

	params, err := form.GetJsonForm[form.Mkdir](c)
	if err != nil {
		response.Error(c, errors.InvalidParameter().Error())
		return
	}

	if err := dirHelper.CheckName(params.Name); err != nil {
		response.Error(c, err.Error())
		return
	}

	userFile := &models.UserFile{
		UserId:   userId,
		ParentId: params.ParentId,
		IsDir:    models.IsDirYes,
		Name:     params.Name,
	}

	if err := userFile.Mkdir(); err != nil {
		response.Error(c, errors.New("创建文件夹失败，原因: "+err.Error()).Error())
		return
	}

	response.Success(c, "创建成功", dto.LoadMkdir(userFile))
}
