package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/dto"
)

// SingleUpload 单文件上传
func SingleUpload(c *gin.Context, userId uint, targetId int) (*dto.SingleUploadDto, error) {

	// 校验参数
	file, err := c.FormFile("file")
	if err != nil {
		return nil, errors.InvalidParameter()
	}

	// 校验文件夹归属
	if err := (&models.UserFile{ParentId: uint(targetId), UserId: userId}).CheckParentId(); err != nil {
		return nil, errors.InvalidParameter()
	}

	return saveSingleFile(c, userId, targetId, file)
}
