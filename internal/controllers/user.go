package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/pkg/jwt"
	"github.com/xstnet/starfire-cloud/pkg/response"
	"regexp"
)

// Login 登录
func Login(c *gin.Context) {
	var loginForm form.LoginForm
	c.ShouldBindJSON(&loginForm)

	user, err := loginForm.Login()
	// 登录失败
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	// 登录成功后下发Token和用户信息
	tokenString, _ := jwt.GenerateToken(user.ID)
	data := gin.H{"token": tokenString, "profile": user.ToDetail()}

	response.Success(c, "登录成功", &data)
}

// Register 用户注册
func Register(c *gin.Context) {
	// 校验参数
	var regForm form.RegisterForm
	c.ShouldBindJSON(&regForm)
	err := regForm.CheckParams()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	// 注册
	user := &models.User{
		Username: regForm.Username,
		Nickname: regForm.Nickname,
		Email:    regForm.Email,
		Password: regForm.Password,
	}
	if err := user.Register(); err != nil {
		response.Error(c, "注册失败, 原因："+err.Error())
		return
	}

	// 注册成功直接下发token和用户信息，不需要进行登录
	tokenString, _ := jwt.GenerateToken(user.ID)
	data := gin.H{"token": tokenString, "profile": user.ToDetail()}

	response.Success(c, "注册成功", &data)
}

// GetProfile 获取用户信息
func GetProfile(c *gin.Context) {
	user := &models.User{}
	if err := user.GetUserById(c.GetUint("userId")); err != nil {
		response.Error(c, "用户不存在")
		return
	}
	response.OkWithData(c, gin.H{"profile": user.ToDetail()})
}

func UpdateProfile(c *gin.Context) {}

// ChangePassword 用户修改密码
func ChangePassword(c *gin.Context) {
	user := &models.User{}

	if err := user.GetUserById(c.GetUint("userId")); err != nil {
		response.Error(c, "用户不存在")
		return
	}

	sourcePassword := c.GetString("source_password")
	if !user.ComparePasswords(sourcePassword) {
		response.Error(c, "原密码错误")
		return
	}

	if c.GetString("password") != c.GetString("password_repeat") {
		response.Error(c, "两次密码输入不一致")
		return
	}

	reg := regexp.MustCompile(`^\S{5,30}$`)
	if !reg.MatchString(c.GetString("password")) {
		response.Error(c, "密码不能包含空格, 且长度在5-20位之间")
		return
	}
	user.Password = c.GetString("password")
	err := user.ChangePassword()
	if err != nil {
		response.Error(c, "修改密码出错，请稍后重试")
		return
	}

	response.Success(c, "ok", "")
}
func ChangeAvatar(c *gin.Context) {}
