package main

import (
	"context"
	"embed"
	"github.com/go-resty/resty/v2"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/manager"
	"github.com/simonkimi/minebangumi/internal/app/router"
	"github.com/simonkimi/minebangumi/internal/app/service"
)

//go:embed dist
var frontendFS embed.FS

func main() {
	manager.Setup()
	mgr := manager.GetInstance()
	conf := mgr.Config

	engine := router.NewRouter(router.NewConfig(&frontendFS, func() string {
		return conf.GetString(config.UserApiToken)
	}, func() *resty.Client {
		return mgr.HttpX.GetTempClient()
	}))

	api := service.NewHttpService(service.NewHttpServiceConfig(
		conf.GetString(config.ServerHost),
		conf.GetInt(config.ServerPort),
		engine,
	))

	exit := make(chan int)

	go func() {
		api.StartHttpService(context.TODO())
		exit <- 1
	}()

	<-exit
}
