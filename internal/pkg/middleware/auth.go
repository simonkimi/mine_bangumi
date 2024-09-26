package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/api"
	"strings"
)

const IsLoginKey = "is_login"

func TokenAuthMiddleware(apiToken func() string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.Next()
			return
		}
		tokenStr = strings.TrimPrefix(tokenStr, "Token ")
		if apiToken() != tokenStr {
			c.Next()
			return
		}

		c.Set(IsLoginKey, true)
		c.Next()
	}
}

func RequireAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetBool(IsLoginKey) {
			c.Next()
			return
		}
		_ = c.Error(api.NewUnAuthError())
		c.Abort()
		return
	}
}

func IsLogin(c *gin.Context) bool {
	return c.GetBool(IsLoginKey)
}
