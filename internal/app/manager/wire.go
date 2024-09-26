//go:build wireinject

package manager

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/repository"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/internal/pkg/database"
	"github.com/simonkimi/minebangumi/pkg/httpx"
	"github.com/simonkimi/minebangumi/pkg/mikan"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
)

func ProvideHttpXConfig(conf *config.Config) *httpx.Config {
	return &httpx.Config{
		ProxyEnabled:  conf.GetBool(config.ProxyEnabled),
		ProxyScheme:   conf.GetString(config.ProxyScheme),
		ProxyHost:     conf.GetString(config.ProxyHost),
		ProxyPort:     conf.GetString(config.ProxyPort),
		ProxyUsername: conf.GetString(config.ProxyUsername),
		ProxyPassword: conf.GetString(config.ProxyPassword),
	}
}

func ProvideTempClient(hx *httpx.HttpX) func() *resty.Client {
	return hx.GetTempClient
}

func ProvideTmdbConfig(conf *config.Config, hx *httpx.HttpX) *tmdb.Config {
	return tmdb.NewConfig(conf.GetString(config.TmdbApiKey), hx.GetTempClient)
}

func InitializeManager() (*Manager, error) {
	wire.Build(
		config.NewConfig,
		ProvideHttpXConfig,
		httpx.NewHttpX,
		ProvideTempClient,
		mikan.NewClient,
		ProvideTmdbConfig,
		tmdb.NewTmdb,
		service.NewScraperService,
		service.NewSourceService,
		database.NewDb,
		repository.NewRepo,
		newManager,
	)
	return nil, nil
}
