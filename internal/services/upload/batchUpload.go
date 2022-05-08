package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"strconv"
)

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
		_, err := saveSingleFile(c, userId, targetId, file)
		if err != nil {
			return err
		}
	}

	return nil
}
