package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/api"
	"net/http"
)

func ResponseWrapperMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			var apiError *api.Error
			if errors.As(err, &apiError) {
				c.JSON(http.StatusOK, gin.H{
					"code":       apiError.Code,
					"message":    apiError.Message,
					"extensions": apiError.Extensions,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":       api.APIStatusEnumInternalServerError,
					"message":    "Internal server error",
					"extensions": gin.H{},
				})
			}
			return
		}
		if c.Writer.Status() == http.StatusNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code":       api.APIStatusEnumNotFound,
				"message":    "404 Not Found",
				"extensions": gin.H{},
			})
			return
		}
	}
}
