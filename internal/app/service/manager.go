package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/repository"
	"github.com/simonkimi/minebangumi/pkg/logger"
	"github.com/simonkimi/minebangumi/pkg/mikan"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
)

type Manager struct {
	Config   *config.Service
	Mikan    *mikan.Client
	HttpX    *HttpX
	Tmdb     *tmdb.Tmdb
	Repo     *repository.Repo
	Scraper  *ScraperService
	Source   *SourceService
	ApiProxy *ApiProxyService
}

func newManager(
	config *config.Service,
	httpX *HttpX,
	mikan *mikan.Client,
	tmdb *tmdb.Tmdb,
	repo *repository.Repo,
	scraper *ScraperService,
	source *SourceService,
	apiProxy *ApiProxyService,
) *Manager {
	return &Manager{
		Config:   config,
		HttpX:    httpX,
		Mikan:    mikan,
		Tmdb:     tmdb,
		Repo:     repo,
		Scraper:  scraper,
		Source:   source,
		ApiProxy: apiProxy,
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
		panic(errors.Wrap(err, "Manager setup failed"))
	}
	instance = i
	logger.Setup()
}
