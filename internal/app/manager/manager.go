package manager

import (
	"fmt"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/pkg/httpx"
	"github.com/simonkimi/minebangumi/pkg/logger"
	"github.com/simonkimi/minebangumi/pkg/mikan"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
	"github.com/sirupsen/logrus"
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
	conf, err := config.Setup()
	if err != nil {
		panic(err)
	}
	logger.Setup()

	instance = &Manager{}
	instance.Config = conf
	instance.HttpX = httpx.NewHttpX(conf.GetBool(config.ProxyEnabled), conf.GetString(config.ProxyScheme), conf.GetString(config.ProxyHost), conf.GetString(config.ProxyPort), conf.GetString(config.ProxyUsername), conf.GetString(config.ProxyPassword))
	instance.Mikan = mikan.NewClient(instance.HttpX.GetTempClient)
	instance.Tmdb = tmdb.NewTmdb(conf.GetString(config.TmdbApiKey), instance.HttpX.GetTempClient)
	instance.Scraper = service.NewScraperService(instance.Tmdb)
	instance.Source = service.NewSourceService(instance.Mikan)
	logrus.Debugf("Manager setup complete")
}
