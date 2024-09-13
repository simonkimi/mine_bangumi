package router

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/graph"
	handler2 "github.com/simonkimi/minebangumi/internal/app/handler"
)
import "github.com/99designs/gqlgen/graphql/handler"

func InitRouter() *gin.Engine {
	r := gin.New()

	graphSrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &handler2.Resolver{},
	}))

	graphqlGroup := r.Group("/graphql")
	{
		graphqlGroup.POST("/query", func(c *gin.Context) {
			graphSrv.ServeHTTP(c.Writer, c.Request)
		})
		graphqlGroup.GET("/", func(c *gin.Context) {
			playground.Handler("GraphQL playground", "/graphql/query").ServeHTTP(c.Writer, c.Request)
		})
	}

	return r
}
