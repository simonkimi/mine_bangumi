package manager

import (
	"fmt"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/pkg/httpx"
	"github.com/simonkimi/minebangumi/pkg/logger"
	"github.com/simonkimi/minebangumi/pkg/mikan"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
)

type Manager struct {
	Config *config.Config
	Mikan  *mikan.Client
	HttpX  *httpx.HttpX
	Tmdb   *tmdb.Tmdb

	Scraper *service.ScraperService
	Source  *service.SourceService
}

var instance *Manager

func GetInstance() *Manager {
	if instance == nil {
		panic(fmt.Errorf("manager is not setup"))
	}
	return instance
}

func Setup() {
	if instance != nil {
		panic(fmt.Errorf("manager is already setup"))
	}
	c, err := config.Setup()
	if err != nil {
		panic(err)
	}

	logger.Setup()

	instance = &Manager{
		Config: c,
	}
}
