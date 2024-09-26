package main

import (
	"context"
	"embed"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/router"
	"github.com/simonkimi/minebangumi/internal/app/service"
)

//go:embed dist
var frontendFS embed.FS

func main() {
	service.Setup()
	mgr := service.GetInstance()

	host := mgr.Config.GetString(config.ServerHost)
	port := mgr.Config.GetInt(config.ServerPort)
	engine := router.NewRouter(router.NewConfig(&frontendFS, mgr.Config, mgr.HttpX))
	api := service.NewHttpService(service.NewHttpServiceConfig(host, port, engine))

	exit := make(chan int)

	go func() {
		api.StartHttpService(context.TODO())
		exit <- 1
	}()

	<-exit
}
