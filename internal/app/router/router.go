package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	_ "github.com/simonkimi/minebangumi/docs"
	"github.com/simonkimi/minebangumi/internal/app/handler"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/fs"
	"net/http"
)

func InitRouter(frontendFs *embed.FS) *gin.Engine {
	r := gin.New()
	middleware.Setup(r)

	if frontendFs != nil {
		r.StaticFileFS("/favicon.ico", "/dist/favicon.ico", http.FS(frontendFs))
		assetsFs, _ := fs.Sub(frontendFs, "dist/assets")
		r.StaticFS("/assets", http.FS(assetsFs))
		r.NoRoute(func(context *gin.Context) {
			data, err := frontendFs.ReadFile("dist/index.html")
			if err != nil {
				context.String(http.StatusInternalServerError, "Error reading index.html")
				return
			}
			context.Data(http.StatusOK, "text/html; charset=utf-8", data)
		})
	}

	apiV1Group(r)
	return r
}

func apiV1Group(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")

	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	configGroup := apiV1.Group("/config")
	{
		configGroup.GET("/system", handler.GetSystem)
		configGroup.POST("/init_user", handler.PostInitUser)
		configGroup.GET("/downloader", handler.GetDownloader)
	}
	userGroup := apiV1.Group("/user")
	{
		userGroup.POST("/login", handler.Login)
	}
	sourceGroup := apiV1.Group("/source")
	{
		sourceGroup.POST("/parse", handler.Source)
		sourceGroup.POST("/scrape", handler.Scrape)
	}
	proxyGroup := apiV1.Group("/proxy")
	{
		proxyGroup.GET("/poster", handler.Poster)
	}
}
