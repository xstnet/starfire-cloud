package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/dto"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/internal/services/user"
)

// Instant 秒传
func Instant(c *gin.Context, userId uint) (*dto.SingleUploadDto, error) {
	params, err := form.GetForm[form.Instant](c)
	if err != nil {
		return nil, errors.InvalidParameter()
	}

	userModel := &models.User{}
	if err := userModel.GetUserById(userId); err != nil {
		return nil, errors.New("用户不存在")
	}

	// 校验文件夹归属
	if err := (&models.UserFile{ParentId: params.TargetId, UserId: userId}).CheckParentId(); err != nil {
		return nil, errors.InvalidParameter()
	}

	// 检查文件是否已存在
	fileModel := &models.File{}
	if ok := fileModel.GetFileByMd5(params.Md5); !ok {
		return nil, errors.New("文件不存在")
	}

	// 检查余量
	if err := user.CheckRemainSpace(userModel, fileModel.Size); err != nil {
		return nil, err
	}

	fileModel.IncRef()
	return bindUserFile(userModel, fileModel, params.TargetId, params.Name)
}
