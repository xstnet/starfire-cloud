package filemanager

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
	"github.com/xstnet/starfire-cloud/pkg/helper/dirHelper"
	"github.com/xstnet/starfire-cloud/pkg/helper/fileHelper"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

// Rename 重命名
func Rename(c *gin.Context) {
	userFile, err := doRename(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	// 需要返回名称，可能会有同名文件导致重命名，因此前端需要使用后端最新的数据
	response.Success(c, "重命名成功", &d.StringMap{"name": userFile.Name})
}

func doRename(c *gin.Context, userId uint) (*models.UserFile, error) {
	params, err := form.GetJsonForm[form.Rename](c)
	if err != nil {
		return nil, errors.InvalidParameter()
	}

	userFileModel := new(models.UserFile)
	if result := userFileModel.DB().First(userFileModel, params.FileId); result.Error != nil || userFileModel.UserId != userId {
		return nil, errors.New("操作对象不存在")
	}

	if userFileModel.IsDir == models.IsDirNo {
		if err := fileHelper.CheckName(params.NewName); err != nil {
			return nil, err
		}
	} else {
		if err := dirHelper.CheckName(params.NewName); err != nil {
			return nil, err
		}
	}

	if err := userFileModel.Rename(params.NewName); err != nil {
		return nil, err
	}

	return userFileModel, nil
}
