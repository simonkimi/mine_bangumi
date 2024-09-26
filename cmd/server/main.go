package main

import (
	"embed"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/router"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/internal/pkg/database"
	"github.com/simonkimi/minebangumi/pkg/logger"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

//go:embed dist
var frontendFS embed.FS

func init() {
	logger.Setup()
	config.NewConfig()
	database.Setup()
}

func main() {
	server := service.GetServerManager()
	engine := router.InitRouter(&frontendFS)

	server.RegisterGin(engine)

	logrus.Warn("Main Starting server...")
	server.StartServer()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Warn("Main Shutting down server...")
	server.Shutdown()
}
