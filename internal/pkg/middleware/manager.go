package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"regexp"
)

var localhostReg = regexp.MustCompile(`^https?://localhost(:\d+)?$`)

func Apply(g *gin.Engine, token func() string) {
	g.Use(gzip.Gzip(gzip.DefaultCompression))
	g.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return localhostReg.MatchString(origin)
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	g.Use(LogrusMiddleware())
	g.Use(ResponseWrapperMiddleware())
	g.Use(TokenAuthMiddleware(token))
}
