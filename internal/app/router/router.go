package router

import (
	graphHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/simonkimi/minebangumi/docs"
	"github.com/simonkimi/minebangumi/internal/app/graph"
	"github.com/simonkimi/minebangumi/internal/app/handler"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	middleware.Setup(r)

	graphSrv := graphHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &handler.Resolver{},
	}))

	v1 := r.Group("/v1")
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	graphqlGroup := v1.Group("/graphql")
	{
		graphqlGroup.POST("/query", func(c *gin.Context) {
			graphSrv.ServeHTTP(c.Writer, c.Request)
		})
		graphqlGroup.GET("/", func(c *gin.Context) {
			playground.Handler("GraphQL playground", "/graphql/query").ServeHTTP(c.Writer, c.Request)
		})
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
	return r
}
