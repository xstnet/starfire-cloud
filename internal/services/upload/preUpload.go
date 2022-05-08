package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/internal/services/user"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
)

// PreUpload 上传前的一些检查操作
func PreUpload(c *gin.Context, userId uint) (*d.StringMap, error) {
	dataForm := form.PreUpload{}
	if err := c.ShouldBindJSON(&dataForm); err != nil {
		return nil, errors.New(err.Error())
	}

	userModel := new(models.User)
	if err := userModel.GetUserById(userId); err != nil {
		return nil, errors.New("用户不存在")
	}

	// 检查余量
	if err := user.CheckRemainSpace(userModel, dataForm.Size); err != nil {
		return nil, err
	}

	// 检查文件是否已存在
	var exist = 0
	fileModel := &models.File{}
	if ok := fileModel.GetFileByMd5(dataForm.Md5); ok {
		exist = 1
	}

	return &d.StringMap{
		"exist": exist,
	}, nil

}
