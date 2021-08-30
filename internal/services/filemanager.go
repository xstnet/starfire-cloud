package services

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/models"
)

func Mkdir(c *gin.Context, userId uint) (*models.UserFile, error) {
	userFile := &models.UserFile{}
	err := c.ShouldBindJSON(&userFile)
	if err != nil {
		return nil, errors.New("参数错误")
	}

	if userFile.Name == "" {
		return nil, errors.New("文件夹名称不能为空")
	}
	matched, err := regexp.MatchString(`^[^/\\\\:\\*\\?\\<\\>\\|\"]{1,255}$`, userFile.Name)
	if err != nil || !matched {
		return nil, errors.New(`文件夹名称不能包含\/:*?"<>|`)
	}

	userFile.UserId = userId

	if err := userFile.Mkdir(); err != nil {
		return nil, errors.New("创建文件夹失败，原因: " + err.Error())
	}

	return userFile, nil
}
