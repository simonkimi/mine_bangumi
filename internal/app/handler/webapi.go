package handler

import (
	"embed"
	grhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/fs"
	"net/http"
)

type WebApi struct {
	Engine     *gin.Engine
	mgr        service.Manager
	frontendFs *embed.FS
}

type WebApiConfig struct {
	frontendFs *embed.FS
	mgr        service.Manager
}

func NewWebApiConfig(frontendFs *embed.FS, mgr service.Manager) *WebApiConfig {
	return &WebApiConfig{
		frontendFs: frontendFs,
		mgr:        mgr,
	}
}

func NewWebApi(conf *WebApiConfig) *WebApi {
	r := gin.New()
	webapi := &WebApi{Engine: r, mgr: conf.mgr, frontendFs: conf.frontendFs}

	middleware.Apply(r, func() string {
		return conf.mgr.GetConfig().GetString(config.UserApiToken)
	})
	webapi.frontend()
	webapi.apiV1Group()

	return webapi
}

func (w *WebApi) frontend() {
	if w.frontendFs == nil {
		return
	}
	r := w.Engine
	r.StaticFileFS("/favicon.ico", "/dist/favicon.ico", http.FS(w.frontendFs))
	assetsFs, _ := fs.Sub(w.frontendFs, "dist/assets")
	r.StaticFS("/assets", http.FS(assetsFs))
	r.NoRoute(func(context *gin.Context) {
		data, err := w.frontendFs.ReadFile("dist/index.html")
		if err != nil {
			context.String(http.StatusInternalServerError, "Error reading index.html")
			return
		}
		context.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
}

func (w *WebApi) apiV1Group() {
	r := w.Engine
	apiV1 := r.Group("/api/v1")
	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := grhandler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: newResolver(w.mgr)}))
	apiV1.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})
	apiV1.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	proxyGroup := apiV1.Group("/proxy")
	{
		proxyGroup.GET("/poster", w.poster)
	}
}
