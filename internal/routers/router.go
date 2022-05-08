package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/controller"
	"github.com/xstnet/starfire-cloud/internal/controller/filemanager"
	"github.com/xstnet/starfire-cloud/internal/controller/recycle"
	"github.com/xstnet/starfire-cloud/internal/controller/upload"
	"github.com/xstnet/starfire-cloud/internal/middleware"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

// todo panic时对前端返回系统错误

func SetupRouters() *gin.Engine {
	r := gin.New()
	r.Use(middleware.RequestCostHandler(), middleware.CorsHandler(), gin.Logger(), gin.Recovery())

	// r.Use(middleware.TokenHandler())

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, d.StringMap{
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
		v1.POST("/login", controller.Login)
		v1.POST("/register", controller.Register)

		// 用户相关
		userGroup := v1.Group("/user", middleware.TokenValidateHandler())
		{
			userGroup.POST("/change-password", controller.ChangePassword)
			userGroup.POST("/change-avatar", controller.ChangeAvatar)
			userGroup.POST("/profile", controller.UpdateProfile)
			userGroup.GET("/profile", controller.GetProfile)
		}

		// 文件(夹)管理
		filemanagerGroup := v1.Group("/filemanager", middleware.TokenValidateHandler())
		{
			filemanagerGroup.POST("/mkdir", filemanager.Mkdir)
			filemanagerGroup.POST("/rename", filemanager.Rename)
			filemanagerGroup.POST("/move", filemanager.Move)
			filemanagerGroup.POST("/copy", filemanager.Copy)
			filemanagerGroup.POST("/delete", filemanager.Delete)
			filemanagerGroup.GET("/list", filemanager.List)
			filemanagerGroup.GET("/dir-list", filemanager.DirList)
		}

		// 回收站操作
		recycleGroup := v1.Group("/recycle", middleware.TokenValidateHandler())
		{
			recycleGroup.GET("/list", recycle.List)
			recycleGroup.POST("/delete", recycle.Delete)
			recycleGroup.POST("/restore", recycle.Restore)
			recycleGroup.POST("/clear", recycle.Clear)
		}

		// 文件上传
		uploadGroup := v1.Group("/upload", middleware.TokenValidateHandler())
		{
			uploadGroup.POST("/batch", upload.BatchUpload)
			uploadGroup.POST("/single-upload", upload.SingleUpload)
			uploadGroup.POST("/pre-upload", upload.PreUpload)
			// 秒传api
			uploadGroup.POST("/instant", upload.Instant)
		}

		// 分享
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

	}

	r.NoRoute(func(c *gin.Context) {
		response.JSON(c, 404, "404", nil)
	})

	return r
}
