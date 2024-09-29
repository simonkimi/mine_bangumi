package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/pkg/database"
	"strings"
)

const IsLoginKey = "is_login"

func TokenAuthMiddleware(conf config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		systemApi := conf.GetString(config.UserApiToken)
		if systemApi == "" {
			c.Set(IsLoginKey, true)
		} else {
			tokenStr := strings.TrimPrefix(c.GetHeader("Authorization"), "Token ")
			if tokenStr == systemApi {
				c.Set(IsLoginKey, true)
			}
		}
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

func DatabaseMigrateMiddleware(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		if db.NeedMigrate() {
			_ = c.Error(&api.Error{
				Message: "Database need migrate",
				Code:    api.APIStatusEnumDatabaseMigrationError,
				Extensions: map[string]any{
					"version":     db.GetSchemaVersion(),
					"app_version": db.GetAppSchemaVersion(),
				},
			})
			c.Abort()
		}
		c.Next()
	}
}

func IsLogin(c *gin.Context) bool {
	return c.GetBool(IsLoginKey)
}
