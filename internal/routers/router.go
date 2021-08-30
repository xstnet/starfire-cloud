package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/controllers"
	"github.com/xstnet/starfire-cloud/internal/middleware"
)

func SetupRouters() *gin.Engine {
	r := gin.New()
	r.Use(middleware.RequestCostHandler(), gin.Logger(), gin.Recovery())

	// r.Use(middleware.TokenHandler())

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	r.GET("/test", middleware.TokenAuthHandler(), func(c *gin.Context) {
		fmt.Println("tttttttt")
	})

	// version 1
	v1 := r.Group("/v1")
	{
		// share := v1.Group("/share")
		// {
		// 	share.POST("/cancle", controllers.Cancle)
		// }
		// /share/cancle
		// /share/list
		// /share/create
		// /share/delete
		// /share/update

		// /s/:name // 分享

		// login & register
		v1.POST("/login", controllers.Login)
		v1.POST("/register", controllers.Register)

		// User
		user := v1.Group("/user", middleware.TokenAuthHandler())
		{
			user.POST("/change-password", controllers.ChangePassword)
			user.POST("/change-avatar", controllers.ChangeAvatar)
			user.POST("/profile", controllers.UpdateProfile)
			user.GET("/profile", controllers.GetProfile)
		}

		// File operation
		file := v1.Group("/file", middleware.TokenAuthHandler())
		{
			file.POST("/mkdir", controllers.Mkdir)
		}

		// 上传
		// upload := v1.Group("/upload", middleware.TokenAuthHandler())
		// {
		// 	upload.POST("/single-file", middleware.TokenAuthHandler())
		// }

	}

	return r
}