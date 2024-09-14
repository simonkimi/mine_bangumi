package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Setup(g *gin.Engine) {
	g.Use(gzip.Gzip(gzip.DefaultCompression))
	g.Use(LogrusMiddleware())
	g.Use(JwtAuthMiddleware())
	g.Use(ResponseWrapperMiddleware())
}
