package config

import (
	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"reflect"
)

var config *appConfig

var configPath = ""

var defaultInit map[string]any

func SetDefaultInit(data map[string]any) {
	defaultInit = data
}

func Setup() {
	wd, err := os.Getwd()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get working directory")
	}
	configPath = filepath.Join(wd, "config.toml")

	viper.SetConfigName("config")
	viper.AddConfigPath(wd)
	viper.SetConfigType("toml")

	config = &appConfig{
		IsNewSystem: false,
		keys:        newConfigKeys(),
	}

	if defaultInit != nil {
		for k, v := range defaultInit {
			viper.SetDefault(k, v)
		}
	}

	err = viper.ReadInConfig()
	var configFileNotFound viper.ConfigFileNotFoundError
	if errors.As(err, &configFileNotFound) {
		config.IsNewSystem = true
		logrus.Warn("Config file not found, use default values")
	}
}

type configKeys struct {
	userUsername *configItem
	userPassword *configItem

	serverIpv4Host *configItem
	serverIpv4Port *configItem
	serverIpv6Host *configItem
	serverIpv6Port *configItem

	downloaderClient *configItem

	qBittorrentApi      *configItem
	qBittorrentUser     *configItem
	qBittorrentPassword *configItem

	aria2Api   *configItem
	aria2Token *configItem

	tmdbApiKey *configItem
}

func newConfigKeys() *configKeys {
	model := &configKeys{
		userUsername:        newConfigItem("user.username", "MBG_USERNAME", "admin"),
		userPassword:        newConfigItem("user.password", "MBG_PASSWORD", "admin"),
		serverIpv4Host:      newConfigItem("server.ipv4_host", "MBG_IPV4_HOST", "0.0.0.0"),
		serverIpv4Port:      newConfigItem("server.ipv4_port", "MBG_IPV4_PORT", "7962"),
		serverIpv6Host:      newConfigItem("server.ipv6_host", "MBG_IPV6_HOST", ""),
		serverIpv6Port:      newConfigItem("server.ipv6_port", "MBG_IPV6_PORT", ""),
		downloaderClient:    newConfigItem("downloader.client", "MBG_DOWNLOADER_CLIENT", ""),
		qBittorrentApi:      newConfigItem("downloader.qBittorrent.api", "MBG_QBITTORRENT_API", "http://127:0.0.1:28080"),
		qBittorrentUser:     newConfigItem("downloader.qBittorrent.username", "MBG_QBITTORRENT_USERNAME", ""),
		qBittorrentPassword: newConfigItem("downloader.qBittorrent.password", "MBG_QBITTORRENT_PASSWORD", ""),
		aria2Api:            newConfigItem("downloader.aria2.api", "MBG_ARIA2_API", "http://localhost:6800/jsonrpc"),
		aria2Token:          newConfigItem("downloader.aria2.token", "MBG_ARIA2_TOKEN", ""),
		tmdbApiKey:          newConfigItem("tmdb.api_key", "MBG_TMDB_API_KEY", ""),
	}
	v := reflect.ValueOf(config).Elem()
	for i := 0; i < v.NumField(); i++ {
		item := v.Field(i).Interface().(*configItem)
		item.register()
	}
	return model
}

func saveConfig() {
	if err := viper.WriteConfigAs(configPath); err != nil {
		logrus.WithError(err).Fatal("Failed to write config file")
	}
}
