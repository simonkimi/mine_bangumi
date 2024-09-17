package middleware

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.AppConfig.User.Password == "" {
			c.Set("claims", &api.UserClaims{Username: config.AppConfig.User.Username})
			c.Next()
			return
		}

		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.Next()
			return
		}
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		claims := &api.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return base64.URLEncoding.DecodeString(config.AppConfig.System.SecretKey)
		})
		if err != nil || !token.Valid {
			c.Next()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func RequireAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if GetClaims(c) == nil {
			_ = c.Error(errno.NewApiError(errno.Unauthorized))
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetClaims(c *gin.Context) *api.UserClaims {
	claims, exist := c.Get("claims")
	if !exist {
		return nil
	}
	return claims.(*api.UserClaims)
}
