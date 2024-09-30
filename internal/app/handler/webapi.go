package handler

import (
	"embed"
	grhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/simonkimi/minebangumi/docs"
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

	middleware.Apply(r)
	r.Use(TokenAuthMiddleware(webapi.mgr.GetConfig()))
	r.Use(DatabaseMigrateMiddleware(webapi.mgr.GetDatabase()))

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
	v1 := w.Engine.Group("/api/v1")
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.POST("/user/login", w.login)
	v1.POST("/user/init", w.initUser)
	v1.GET("/system/status", w.systemStatus)

	loginV1 := v1.Group("")
	{
		loginV1.Use(RequireAuthMiddleware())
		srv := grhandler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: newResolver(w.mgr)}))
		loginV1.POST("/graph", func(c *gin.Context) {
			srv.ServeHTTP(c.Writer, c.Request)
		})
		loginV1.GET("/", func(c *gin.Context) {
			playground.Handler("GraphQL playground", "/playground").ServeHTTP(c.Writer, c.Request)
		})

		proxyGroup := loginV1.Group("/proxy")
		{
			proxyGroup.GET("/poster", w.poster)
		}
	}
}
