package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/internal/utils"
)

// 登录
func Login(c *gin.Context) {
	var loginForm form.LoginForm
	c.ShouldBindJSON(&loginForm)

	user, err := loginForm.Login()
	// 登录失败
	if err != nil {
		c.JSON(200, utils.ResponseError(err.Error()))
		return
	}

	// 登录成功后下发Token和用户信息
	tokenString, _ := utils.GenerateToken(int(user.ID))
	data := gin.H{"token": tokenString, "profile": user.ToDetail()}

	c.JSON(200, utils.ResponseSuccess("登录成功", &data))
}

// 用户注册
func Register(c *gin.Context) {
	// 校验参数
	var regForm form.RegisterForm
	c.ShouldBindJSON(&regForm)
	err := regForm.CheckParams()
	if err != nil {
		c.JSON(200, utils.ResponseError(err.Error()))
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
		c.JSON(200, utils.ResponseError("注册失败, 原因："+err.Error()))
		return
	}

	// 注册成功直接下发token和用户信息，不需要进行登录
	tokenString, _ := utils.GenerateToken(int(user.ID))
	data := gin.H{"token": tokenString, "profile": user.ToDetail()}

	c.JSON(200, utils.ResponseSuccess("注册成功", &data))
}
func GetProfile(c *gin.Context)     {}
func UpdateProfile(c *gin.Context)  {}
func ChangePassword(c *gin.Context) {}
func ChangeAvatar(c *gin.Context)   {}
