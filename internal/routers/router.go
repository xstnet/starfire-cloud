package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/controllers"
	"github.com/xstnet/starfire-cloud/internal/middleware"
	"github.com/xstnet/starfire-cloud/pkg/response"
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
		filemanager := v1.Group("/filemanager", middleware.TokenAuthHandler())
		{
			filemanager.POST("/mkdir", controllers.Mkdir)
			filemanager.POST("/rename", controllers.Rename)
			filemanager.POST("/move", controllers.Move)
			filemanager.POST("/copy", controllers.Copy)
			filemanager.POST("/delete", controllers.Delete)
			filemanager.GET("/list", controllers.List)
		}

		// File operation
		recycle := v1.Group("/recycle", middleware.TokenAuthHandler())
		{
			recycle.GET("/list", controllers.RecycleList)
			recycle.POST("/delete", controllers.RecycleDelete)
			recycle.POST("/restore", controllers.RecycleRestore)
			recycle.POST("/clear", controllers.RecycleClear)
		}

		// 上传
		upload := v1.Group("/upload", middleware.TokenAuthHandler())
		{
			upload.POST("/batch", middleware.TokenAuthHandler(), controllers.BatchUpload)
			upload.POST("/pre-upload", middleware.TokenAuthHandler(), controllers.PreUpload)
		}

	}

	r.NoRoute(func(c *gin.Context) {
		response.Error(c, "无效的路由")
	})

	return r
}
