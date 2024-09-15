package router

import (
	"embed"
	graphHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/simonkimi/minebangumi/docs"
	"github.com/simonkimi/minebangumi/internal/app/graph"
	"github.com/simonkimi/minebangumi/internal/app/handler"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
		r.GET("/", func(context *gin.Context) {
			data, err := frontendFs.ReadFile("dist/index.html")
			if err != nil {
				context.String(http.StatusInternalServerError, "Error reading index.html")
				return
			}
			context.Data(http.StatusOK, "text/html; charset=utf-8", data)
		})
	}

	v1 := r.Group("/v1")
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	graphSrv := graphHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &handler.Resolver{},
	}))
	graphqlGroup := v1.Group("/graphql")
	{
		graphqlGroup.POST("/query", func(c *gin.Context) {
			graphSrv.ServeHTTP(c.Writer, c.Request)
		})
		graphqlGroup.GET("/", func(c *gin.Context) {
			playground.Handler("GraphQL playground", "/graphql/query").ServeHTTP(c.Writer, c.Request)
		})
	}
	configGroup := v1.Group("/config")
	{
		configGroup.GET("/system", handler.System)
		configGroup.POST("/init_user", handler.InitUser)
	}
	userGroup := v1.Group("/user")
	{
		userGroup.POST("/login", handler.Login)
	}
	sourceGroup := v1.Group("/source")
	{
		sourceGroup.POST("/parse", handler.Source)
		sourceGroup.POST("/scrape", handler.Scrape)
	}
	proxyGroup := v1.Group("/proxy")
	{
		proxyGroup.GET("/poster", handler.Poster)
	}
	return r
}
