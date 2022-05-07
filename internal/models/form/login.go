package form

import (
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
)

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 校验登录参数
func (l *LoginForm) Login() (*models.User, error) {
	var user = new(models.User)

	if res := user.DB().Where("username = ?", l.Username).First(user); res.Error != nil {
		return nil, errors.New("账号或密码错误")
	}

	if !user.ComparePasswords(l.Password) {
		return nil, errors.New("账号或密码错误")
	}

	return user, nil
}
