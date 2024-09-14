package router

import (
	graphHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/graph"
	"github.com/simonkimi/minebangumi/internal/app/handler"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	middleware.Setup(r)

	graphSrv := graphHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &handler.Resolver{},
	}))

	v1 := r.Group("/v1")
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
		sourceGroup.POST("/parse", handler.ParseSource)
	}
	return r
}
