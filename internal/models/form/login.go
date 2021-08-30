package form

import (
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 校验登录参数
func (l *LoginForm) Login() (*models.User, error) {
	if l.Username == "" {
		return nil, errors.New("请输入用户名")
	}
	if l.Password == "" {
		return nil, errors.New("请输入密码")
	}

	var user models.User

	if res := user.DB().Where("username = ?", l.Username).First(&user); res.Error != nil {
		return nil, errors.New("用户不存在")
	}

	if !user.ComparePasswords(l.Password) {
		return nil, errors.New("密码错误")
	}

	return &user, nil
}
