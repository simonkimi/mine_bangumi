//go:build server

package main

import (
	"github.com/sirupsen/logrus"
	"minebangumi/pkg/logger"
	"minebangumi/service"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	logger.Setup()
}

func main() {
	server := service.NewServer()
	logrus.Warn("Starting server...")
	server.Start()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Warn("Shutting down server...")
	server.Shutdown()
}
