package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/controllers"
	"github.com/xstnet/starfire-cloud/internal/controllers/upload"
	"github.com/xstnet/starfire-cloud/internal/middleware"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

// todo panic时对前端返回系统错误

func SetupRouters() *gin.Engine {
	r := gin.New()
	r.Use(middleware.RequestCostHandler(), middleware.CorsHandler(), gin.Logger(), gin.Recovery())

	// r.Use(middleware.TokenHandler())

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	r.GET("/login/test", middleware.TokenValidateHandler(), func(c *gin.Context) {
		c.JSON(0, "has login")
	})

	// version 1
	v1 := r.Group("/api/v1")
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
		user := v1.Group("/user", middleware.TokenValidateHandler())
		{
			user.POST("/change-password", controllers.ChangePassword)
			user.POST("/change-avatar", controllers.ChangeAvatar)
			user.POST("/profile", controllers.UpdateProfile)
			user.GET("/profile", controllers.GetProfile)
		}

		// File operation
		filemanager := v1.Group("/filemanager", middleware.TokenValidateHandler())
		{
			filemanager.POST("/mkdir", controllers.Mkdir)
			filemanager.POST("/rename", controllers.Rename)
			filemanager.POST("/move", controllers.Move)
			filemanager.POST("/copy", controllers.Copy)
			filemanager.POST("/delete", controllers.Delete)
			filemanager.GET("/list", controllers.List)
			filemanager.GET("/dir-list", controllers.DirList)
		}

		// Recycle File operation
		recycle := v1.Group("/recycle", middleware.TokenValidateHandler())
		{
			recycle.GET("/list", controllers.RecycleList)
			recycle.POST("/delete", controllers.RecycleDelete)
			recycle.POST("/restore", controllers.RecycleRestore)
			recycle.POST("/clear", controllers.RecycleClear)
		}

		// 上传
		uploadManager := v1.Group("/upload", middleware.TokenValidateHandler())
		{
			uploadManager.POST("/batch", upload.BatchUpload)
			uploadManager.POST("/single-upload", upload.SingleUpload)
			uploadManager.POST("/pre-upload", upload.PreUpload)
			// 秒传api
			uploadManager.POST("/instant", upload.Instant)
		}

	}

	r.NoRoute(func(c *gin.Context) {
		response.Error(c, "无效的路由")
	})

	return r
}
