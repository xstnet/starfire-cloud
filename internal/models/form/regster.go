package form

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/xstnet/starfire-cloud/internal/db"
	"github.com/xstnet/starfire-cloud/internal/models"
)

type RegisterForm struct {
	Username       string `json:"username"`
	Nickname       string `json:"nickname"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	Email          string `json:"email"`
}

func (r *RegisterForm) CheckParams() error {
	fmt.Println("form:", r)
	if r.Username == "" {
		return errors.New("用户名不能为空")
	}
	if r.Nickname == "" {
		return errors.New("昵称不能为空")
	}
	if r.Password == "" {
		return errors.New("密码不能为空")
	}
	if r.Password != r.PasswordRepeat {
		return errors.New("两次密码输入不一致")
	}

	reg := regexp.MustCompile(`^[\w-@\.]{5,50}$`)
	if !reg.MatchString(r.Username) {
		return errors.New("用户名只能使用字母/数字/@_-, 且长度在5-50位之间")
	}

	reg = regexp.MustCompile(`^\S{5,30}$`)
	if !reg.MatchString(r.Password) {
		return errors.New("密码不能包含空格, 且长度在5-20位之间")
	}

	reg = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	if !reg.MatchString(r.Email) {
		return errors.New("邮箱格式不正确")
	}

	r.Nickname = strings.TrimSpace(r.Nickname)
	nicknameLength := utf8.RuneCountInString(r.Nickname)
	if nicknameLength < 1 || nicknameLength > 50 {
		return errors.New("昵称长度在1-50个字符之间")
	}

	reg = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	if !reg.MatchString(r.Email) {
		return errors.New("邮箱格式不正确")
	}

	if res := db.DB.Where("username = ?", r.Username).First(&models.User{}); res.Error == nil {
		return errors.New("用户名已被占用")
	}

	if res := db.DB.Where("email = ?", r.Email).First(&models.User{}); res.Error == nil {
		return errors.New("邮箱已被占用")
	}

	return nil
}
