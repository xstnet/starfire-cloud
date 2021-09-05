package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/configs"
	"github.com/xstnet/starfire-cloud/internal/common"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/pkg/systeminfo"
)

func UploadFile(c *gin.Context, userId uint) error {
	targetId, err := strconv.Atoi(c.PostForm("target_id"))

	if err != nil || targetId < 0 {
		return errors.InvalidParameter()
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	err = saveSingleFile(c, userId, targetId, file)
	fmt.Println(err)

	return err
}

// 暂不使用
func BatchUpload(c *gin.Context, userId uint) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["files"]
	if len(files) < 1 {
		return errors.New("请选择文件")
	}

	targetId, err := strconv.Atoi(c.PostForm("target_id"))
	if err != nil || targetId < 0 {
		return errors.InvalidParameter()
	}

	for _, file := range files {
		err := saveSingleFile(c, userId, targetId, file)
		if err != nil {
			return err
		}
	}

	return nil
}

// 获取上传文件的md5
func getUploadFileMd5(file *multipart.FileHeader) (string, error) {
	fh, err := file.Open()
	if err != nil {
		return "", err
	}

	md5h := md5.New()
	io.Copy(md5h, fh)

	return hex.EncodeToString(md5h.Sum(nil)), nil
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

func saveSingleFile(c *gin.Context, userId uint, targetId int, file *multipart.FileHeader) error {
	// 检查文件名是否合法
	if err := common.CheckFilename(file.Filename); err != nil {
		return errors.InvalidParameter()
	}
	if len(file.Header["Content-Type"]) < 1 || file.Size < 0 {
		return errors.InvalidParameter()
	}

	user := &models.User{}
	if err := user.GetUserById(userId); err != nil {
		return errors.New("用户不存在")
	}
	fmt.Println(user)
	// 检查余量
	if err := checkRemainSpace(user, uint64(file.Size)); err != nil {
		return err
	}

	// 使用文件内容的MD5值做为文件名
	md5Str, err := getUploadFileMd5(file)
	if err != nil {
		return err
	}

	// 检查文件是否已存在
	// 文件不存在则上传
	fileModel := &models.File{}
	if ok := fileModel.GetFileByMd5(md5Str); !ok {
		// 转小写， 并去除面的 .
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(file.Filename)), ".")

		relativePath, err := getAndCreateSavePath(userId)
		if err != nil {
			return err
		}

		fullPath := filepath.Join(relativePath, md5Str+"."+ext)

		// 上传
		err = c.SaveUploadedFile(file, filepath.Join(configs.Upload.UploadRootPath, fullPath))
		if err != nil {
			return err
		}
		// 入库
		err = fileModel.Create(userId, uint64(file.Size), md5Str, fullPath, ext, file.Header["Content-Type"][0])
		if err != nil {
			return err
		}

	} else {
		// 更新引用数量
		fileModel.IncRef()
	}

	// 将文件对象绑定到用户的文件
	userFile := &models.UserFile{
		UserId:   userId,
		FileId:   fileModel.ID,
		Name:     file.Filename,
		ParentId: uint(targetId),
		IsDir:    models.IS_DIR_NO,
	}
	if err := userFile.BindFile(); err != nil {
		return err
	}
	// 更新用户的存储空间
	if err := user.UpdateUsedSpace(fileModel.Size); err != nil {
		return err
	}

	return nil
}

// 检查剩余的存储空间是否足够
func checkRemainSpace(user *models.User, size uint64) error {
	// 判断用户的剩余空间是否足够
	if user.TotalSpace > 0 && user.TotalSpace-user.UsedSpace < size {
		return errors.New("可用上传空间不足，当前剩余: " + common.FormatFileSize(user.TotalSpace-user.UsedSpace))
	}
	// 判断磁盘余量
	diskInfo := systeminfo.DiskInfo(configs.Upload.UploadRootPath)
	if diskInfo.Total == 0 {
		return errors.New("获取磁盘信息失败")
	}
	if diskInfo.Free < size {
		return errors.New("磁盘剩余空间不足，当前余量: " + common.FormatFileSize(diskInfo.Free))
	}
	return nil
}