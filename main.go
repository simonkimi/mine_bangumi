package main

import (
	"context"
	"embed"
	"github.com/simonkimi/minebangumi/internal/app/handler"
	"github.com/simonkimi/minebangumi/internal/app/service"
)

//go:embed frontend/dist
var frontendFS embed.FS

func main() {
	mgr := service.Initialize()

	webapi := handler.NewWebApi(handler.NewWebApiConfig(&frontendFS, mgr))

	exit := make(chan int)
	go func() {
		mgr.GetHttpService().StartHttpService(context.TODO(), webapi.Engine)
		exit <- 1
	}()

	<-exit
}
