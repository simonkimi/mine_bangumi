package gui

import (
	"context"
	"github.com/simonkimi/minebangumi/internal/app/manager"
)

type App struct {
	ctx    context.Context
	server *manager.ServerManager
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

}

//goland:noinspection ALL
func (a *App) DomReady(ctx context.Context) {
	a.server = manager.NewServerManager()
	a.server.Start()
}

//goland:noinspection ALL
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

//goland:noinspection ALL
func (a *App) Shutdown(ctx context.Context) {
	a.server.Shutdown()
}
