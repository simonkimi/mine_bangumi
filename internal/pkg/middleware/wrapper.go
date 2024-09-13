package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"net/http"
)

func ResponseWrapperMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			var apiError errno.ApiError
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
			c.Abort()
			return
		}

		switch c.Writer.Status() {
		case http.StatusNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "404 Not Found",
			})
		case http.StatusOK:
			data, exists := c.Get("data")
			if exists {
				c.JSON(http.StatusOK, gin.H{
					"code": errno.Success,
					"data": data,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    errno.Success,
					"message": nil,
				})
			}
		}
	}
}
