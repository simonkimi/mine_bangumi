package webapi

import (
	"context"
	"embed"
	"errors"
	"fmt"
	grhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/handler"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/fs"
	"net"
	"net/http"
	"sync"
)

type HttpService struct {
	host       string
	ActualPort int
	engine     *gin.Engine
}

type HttpServiceConfig struct {
	Host   string
	Port   int
	engine *gin.Engine
}

func NewHttpServiceConfig(host string, port int, engine *gin.Engine) *HttpServiceConfig {
	return &HttpServiceConfig{Host: host, Port: port, engine: engine}

}

func NewHttpService(config *HttpServiceConfig) *HttpService {
	return &HttpService{host: config.Host, ActualPort: config.Port, engine: config.engine}
}

func (s *HttpService) StartHttpService(ctx context.Context) {
	portChan := make(chan int)
	exitChan := make(chan int)
	go s.startHttpService(ctx, s.engine, portChan, exitChan)
	s.ActualPort = <-portChan
	<-exitChan
}

func (s *HttpService) startHttpService(ctx context.Context, engine *gin.Engine, portChan, exitChan chan int) {
	gin.SetMode(gin.DebugMode)

	var ipv4Listener net.Listener
	var lock sync.Mutex

	addr := fmt.Sprintf("%s:%d", s.host, s.ActualPort)
	go func() {
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			logrus.WithError(err).Fatalf("Failed to listen on %s", addr)
		}
		lock.Lock()
		ipv4Listener = listener
		lock.Unlock()
		port := listener.Addr().(*net.TCPAddr).Port
		portChan <- port
		logrus.Infof("Starting server on %s:%d", s.host, port)

		err = http.Serve(listener, engine)
		if err != nil && !errors.Is(err, net.ErrClosed) {
			logrus.Errorf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	lock.Lock()
	if ipv4Listener != nil {
		if err := ipv4Listener.Close(); err != nil {
			logrus.Errorf("Server shutdown failed: %v", err)
		}
	}
	lock.Unlock()
	exitChan <- 1
	logrus.Println("Server exiting")
}

type Router struct {
	Engine *gin.Engine
}

type Config struct {
	frontendFs *embed.FS
	config     *config.Service
	httpx      *service.HttpX
}

func NewConfig(frontendFs *embed.FS, config *config.Service, httpx *service.HttpX) *Config {
	return &Config{
		frontendFs: frontendFs,
		config:     config,
		httpx:      httpx,
	}
}

func NewRouter(conf *Config) *gin.Engine {
	r := gin.New()
	middleware.Apply(r, func() string {
		return conf.config.GetString(config.UserApiToken)
	})
	frontend(r, conf.frontendFs)
	h := handler.NewHttpHandler(conf.httpx)
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
