package main

import (
	"embed"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/manager"
	"github.com/simonkimi/minebangumi/internal/app/router"
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
	config.Setup()
	database.Setup()
}

func main() {
	server := manager.GetServerManager()
	engine := router.InitRouter(&frontendFS)

	server.RegisterGin(engine)

	logrus.Warn("Starting server...")
	server.StartServer()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Warn("Shutting down server...")
	server.Shutdown()
}
