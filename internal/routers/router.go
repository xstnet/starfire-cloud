package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/controllers"
	"github.com/xstnet/starfire-cloud/internal/middleware"
)

func MyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取header里的token值，没有话，就通过 c.Abort() 方法取消请求的继续进行，从而抛出异常
		if token := c.GetHeader("token"); token == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "token not found"})
			c.Abort()
		} else {
			// 让程序继续正常运行
			c.Next()
		}
	}
}

func SetupRouters() *gin.Engine {
	r := gin.Default()

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	r.GET("/test", middleware.TokenHandler(), func(c *gin.Context) {
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

		// user
		user := v1.Group("/user")
		{
			user.POST("/login", controllers.Login)
			user.POST("/register", controllers.Register)
			user.POST("/change-password", controllers.ChangePassword)
			user.POST("/change-avatar", controllers.ChangeAvatar)
			user.POST("/profile", controllers.UpdateProfile)
			user.GET("/profile", controllers.GetProfile)
		}

		// 上传
		// /upload/file POST

	}

	return r
}
