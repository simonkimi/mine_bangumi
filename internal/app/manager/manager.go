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

func newManager(
	config *config.Config,
	httpX *httpx.HttpX,
	mikan *mikan.Client,
	tmdb *tmdb.Tmdb,
	scraper *service.ScraperService,
	source *service.SourceService,
) *Manager {
	return &Manager{
		Config:  config,
		HttpX:   httpX,
		Mikan:   mikan,
		Tmdb:    tmdb,
		Scraper: scraper,
		Source:  source,
	}
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

	i, err := InitializeManager()
	if err != nil {
		panic(err)
	}
	instance = i
	logger.Setup()
}
