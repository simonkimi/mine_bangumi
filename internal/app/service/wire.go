//go:build wireinject

package service

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/repository"
	"github.com/simonkimi/minebangumi/internal/pkg/database"
	"github.com/simonkimi/minebangumi/pkg/mikan"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
)

func provideHttpXConfig(conf *config.Service) *HttpxConfig {
	return &HttpxConfig{
		ProxyEnabled:  conf.GetBool(config.ProxyEnabled),
		ProxyScheme:   conf.GetString(config.ProxyScheme),
		ProxyHost:     conf.GetString(config.ProxyHost),
		ProxyPort:     conf.GetString(config.ProxyPort),
		ProxyUsername: conf.GetString(config.ProxyUsername),
		ProxyPassword: conf.GetString(config.ProxyPassword),
	}
}

func provideTempClient(hx *HttpX) func() *resty.Client {
	return hx.GetTempClient
}

func provideTmdbConfig(conf *config.Service, hx *HttpX) *tmdb.Config {
	return tmdb.NewConfig(conf.GetString(config.TmdbApiKey), hx.GetTempClient)
}

func InitializeManager() (*Manager, error) {
	wire.Build(
		config.NewConfig,
		provideHttpXConfig,
		NewHttpX,
		provideTempClient,
		mikan.NewClient,
		provideTmdbConfig,
		tmdb.NewTmdb,
		newScraperService,
		newSourceService,
		database.NewDb,
		repository.NewRepo,
		newApiProxyService,
		newManager,
	)
	return nil, nil
}
