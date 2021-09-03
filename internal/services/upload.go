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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/configs"
	"github.com/xstnet/starfire-cloud/internal/errors"
)

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

	now := time.Now()
	// 目录格式 root_path/year/month/day/user_id/md5(file).ext
	relativePath := fmt.Sprintf("%s/%s/%s/%d", now.Format("2006"), now.Format("01"), now.Format("02"), userId)
	rootPath := configs.Upload.UploadRootPath
	// 如果没有配置上传文件的根目录， 默认存储在程序运行目录下的 /uploads下面， 格式同上
	if rootPath == "" {
		currentPath, err := os.Getwd()
		if err != nil {
			return errors.SystemError()
		}

		rootPath = filepath.Join(currentPath, "/uploads")
	}

	savePath := filepath.Join(rootPath, relativePath)
	if err := os.MkdirAll(savePath, 0766); err != nil {
		return errors.New("保存文件目录失败")
	}

	for _, file := range files {

		// 使用文件内容的MD5值做为文件名
		filename, err := getUploadFileMd5(file)
		if err != nil {
			return err
		}

		ext := filepath.Ext(file.Filename)
		if ext != "" {
			filename += ext
		}

		err = c.SaveUploadedFile(file, filepath.Join(savePath, filename))
		if err != nil {
			fmt.Println("upload err:", err)
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
