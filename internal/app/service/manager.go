package service

import (
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/repository"
	"github.com/simonkimi/minebangumi/internal/pkg/database"
	"github.com/simonkimi/minebangumi/pkg/logger"
	"github.com/simonkimi/minebangumi/pkg/mikan"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
)

//go:generate mockery --name=Manager
type Manager interface {
	GetConfig() config.Config
	GetMikan() *mikan.Client
	GetHttpX() *HttpX
	GetTmdb() *tmdb.Tmdb
	GetRepo() *repository.Repo
	GetDatabase() *database.Database
	GetScraper() *ScraperService
	GetSource() *SourceService
	GetApiProxy() *ApiProxyService
	GetHttpService() *HttpService
}

type ManagerImpl struct {
	config      config.Config
	mikan       *mikan.Client
	httpX       *HttpX
	tmdb        *tmdb.Tmdb
	database    *database.Database
	repo        *repository.Repo
	scraper     *ScraperService
	source      *SourceService
	apiProxy    *ApiProxyService
	httpService *HttpService
}

func newManager(
	config config.Config,
	httpX *HttpX,
	mikan *mikan.Client,
	tmdb *tmdb.Tmdb,
	database *database.Database,
	repo *repository.Repo,
	scraper *ScraperService,
	source *SourceService,
	apiProxy *ApiProxyService,
	httpService *HttpService,
) Manager {
	return &ManagerImpl{
		config:      config,
		httpX:       httpX,
		mikan:       mikan,
		tmdb:        tmdb,
		repo:        repo,
		database:    database,
		scraper:     scraper,
		source:      source,
		apiProxy:    apiProxy,
		httpService: httpService,
	}
}

func Initialize() Manager {
	instance, err := InitializeManager()
	if err != nil {
		panic(errors.Wrap(err, "Manager setup failed"))
	}
	logger.Setup()
	return instance
}

func (m *ManagerImpl) GetConfig() config.Config {
	return m.config
}

func (m *ManagerImpl) GetMikan() *mikan.Client {
	return m.mikan
}

func (m *ManagerImpl) GetHttpX() *HttpX {
	return m.httpX
}

func (m *ManagerImpl) GetTmdb() *tmdb.Tmdb {
	return m.tmdb
}

func (m *ManagerImpl) GetRepo() *repository.Repo {
	return m.repo
}

func (m *ManagerImpl) GetScraper() *ScraperService {
	return m.scraper
}

func (m *ManagerImpl) GetSource() *SourceService {
	return m.source
}

func (m *ManagerImpl) GetApiProxy() *ApiProxyService {
	return m.apiProxy
}

func (m *ManagerImpl) GetHttpService() *HttpService {
	return m.httpService
}

func (m *ManagerImpl) GetDatabase() *database.Database {
	return m.database
}
