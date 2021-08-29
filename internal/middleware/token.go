package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/common"
	"github.com/xstnet/starfire-cloud/internal/utils"
)

func TokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if tokenString == "" {
			c.JSON(200, utils.ResponseError("请先登录后再操作"))
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			// 解析Token错误
			c.JSON(200, utils.ResponseError(err.Error()))
			c.Abort()
			return
		}

		// 处理token过期逻辑
		if tokenIsExpired(claims.ExpiresAt) {
			// 不能自动刷新token, 需要重新登录
			if !canRefreshToken(claims.IssuedAt) {
				c.JSON(200, utils.ResponseJSON(common.CODE_RELOGIN, "登录已过期，请重新登录", nil))
				c.Abort()
				return
			}
			// 自动刷新token
			tokenString, err := refreshToken(claims.UserId)
			if err != nil {
				// todo: log
				c.JSON(200, utils.ResponseError("系统错误，请重试"))
				c.Abort()
				return
			}

			// 将新token存储在Header中返回给客户端
			c.Writer.Header().Add("Authorization", tokenString)
		}

		c.Set("userId", claims.UserId)

		c.Next()
	}
}

// token是否已过期
func tokenIsExpired(exp int64) bool {
	return time.Now().Unix() > exp
	// return iat+int64(utils.TokenRemeberDuration) > now
}

// 是否能自动刷新token
func canRefreshToken(iat int64) bool {
	return (iat+int64(utils.TokenRemeberDuration) > time.Now().Unix())
}

// 自动刷新Token
func refreshToken(userId int) (string, error) {
	tokenString, err := utils.GenerateToken(userId)
	return tokenString, err
}
