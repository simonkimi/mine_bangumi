package router

import (
	"embed"
	grhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	_ "github.com/simonkimi/minebangumi/docs"
	"github.com/simonkimi/minebangumi/internal/app/handler"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/fs"
	"net/http"
)

type Router struct {
	Engine *gin.Engine
}

type Config struct {
	frontendFs   *embed.FS
	token        func() string
	getTmpClient func() *resty.Client
}

func NewConfig(frontendFs *embed.FS, token func() string, getTmpClient func() *resty.Client) *Config {
	return &Config{
		frontendFs:   frontendFs,
		token:        token,
		getTmpClient: getTmpClient,
	}
}

func NewRouter(config *Config) *gin.Engine {
	r := gin.New()
	middleware.Apply(r, config.token)
	frontend(r, config.frontendFs)
	h := handler.NewHttpHandler(config.getTmpClient)
	apiV1Group(r, h)

	return r
}

func frontend(r *gin.Engine, frontendFs *embed.FS) {
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

func apiV1Group(r *gin.Engine, h *handler.HttpHandler) {
	apiV1 := r.Group("/api/v1")
	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := grhandler.NewDefaultServer(handler.NewExecutableSchema(handler.Config{Resolvers: &handler.Resolver{}}))
	apiV1.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})
	apiV1.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	proxyGroup := apiV1.Group("/proxy")
	{
		proxyGroup.GET("/poster", h.Poster)
	}
}
