package gui

import (
	"context"
	"embed"
	"github.com/simonkimi/minebangumi/internal/app/router"
	"github.com/simonkimi/minebangumi/internal/app/service"
)

type App struct {
	ctx        context.Context
	server     *service.ServerManager
	frontendFs *embed.FS
}

func NewApp(frontendFs *embed.FS) *App {
	return &App{
		frontendFs: frontendFs,
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	server := service.GetServerManager()
	a.server = server
	server.RegisterGin(router.InitRouter(a.frontendFs))
	server.StartServer()
}

//goland:noinspection ALL
func (a *App) DomReady(ctx context.Context) {

}

//goland:noinspection ALL
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

//goland:noinspection ALL
func (a *App) Shutdown(ctx context.Context) {
	a.server.Shutdown()
}

func (a *App) GetIpv4Port() int {
	return a.server.Ipv4Port
}

func (a *App) GetIpv6Port() int {
	return a.server.Ipv6Port
}
