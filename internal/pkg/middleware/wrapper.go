package middleware

import (
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"net/http"
)

func ResponseWrapperMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			var apiError *errno.ApiError
			if errors.As(err, &apiError) {
				c.JSON(http.StatusOK, gin.H{
					"code":    apiError.Code,
					"message": apiError.Message,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    errno.InternalServerError,
					"message": "Internal server error",
				})
			}
			return
		}
	}
}
