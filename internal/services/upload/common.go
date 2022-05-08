package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/configs"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/dto"
	"github.com/xstnet/starfire-cloud/internal/services/user"
	"github.com/xstnet/starfire-cloud/pkg/helper/crypto"
	"github.com/xstnet/starfire-cloud/pkg/helper/fileHelper"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func saveSingleFile(c *gin.Context, userId uint, targetId int, uploadFile *multipart.FileHeader) (*dto.SingleUploadDto, error) {
	// 检查文件名是否合法
	if err := fileHelper.CheckName(uploadFile.Filename); err != nil {
		return nil, errors.InvalidParameter()
	}
	if len(uploadFile.Header["Content-Type"]) < 1 || uploadFile.Size < 0 {
		return nil, errors.InvalidParameter()
	}

	userModel := &models.User{}
	if err := userModel.GetUserById(userId); err != nil {
		return nil, errors.New("用户不存在")
	}

	// 检查余量
	if err := user.CheckRemainSpace(userModel, uint64(uploadFile.Size)); err != nil {
		return nil, err
	}

	// 使用文件内容的MD5值做为文件名
	fileHandler, err := uploadFile.Open()
	if err != nil {
		return nil, err
	}
	defer fileHandler.Close()

	md5Str := crypto.Md5File(fileHandler)

	// 检查文件是否已存在
	// 文件不存在则上传
	fileModel := &models.File{}
	if ok := fileModel.GetFileByMd5(md5Str); !ok {
		// 转小写， 并去除面的 .
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(uploadFile.Filename)), ".")

		relativePath, err := getAndCreateSavePath(userId)
		if err != nil {
			return nil, err
		}

		fullPath := filepath.Join(relativePath, md5Str+"."+ext)

		// 上传
		err = c.SaveUploadedFile(uploadFile, filepath.Join(configs.Upload.UploadRootPath, fullPath))
		if err != nil {
			return nil, err
		}
		// 入库
		err = fileModel.Create(userId, uint64(uploadFile.Size), md5Str, fullPath, ext, uploadFile.Header["Content-Type"][0])
		if err != nil {
			return nil, err
		}

	} else {
		// 更新引用数量
		fileModel.IncRef()
	}

	return bindUserFile(userModel, fileModel, uint(targetId), uploadFile.Filename)
}

func bindUserFile(userModel *models.User, file *models.File, targetId uint, originName string) (*dto.SingleUploadDto, error) {
	// 将文件对象绑定到用户的文件
	userFile := &models.UserFile{
		UserId:   userModel.ID,
		FileId:   file.ID,
		Name:     originName,
		ParentId: targetId,
		IsDir:    models.IsDirNo,
	}
	if err := userFile.BindFile(); err != nil {
		return nil, err
	}
	// 更新用户的存储空间
	if err := userModel.UpdateUsedSpace(file.Size); err != nil {
		return nil, err
	}

	return dto.LoadSingleUpload(userFile, file), nil
}

// 获取文件保存的目录， 如果不存在， 会自动创建
// 返回的是相对目录， 数据库中不存储绝对目录
func getAndCreateSavePath(userId uint) (relativePath string, err error) {
	now := time.Now()
	// 目录格式 root_path/year/month/day/user_id/md5(file).ext
	relativePath = fmt.Sprintf("%s/%s/%s/%d", now.Format("2006"), now.Format("01"), now.Format("02"), userId)

	fullPath := filepath.Join(configs.Upload.UploadRootPath, relativePath)
	if err = os.MkdirAll(fullPath, 0766); err != nil {
		return
	}

	return
}
