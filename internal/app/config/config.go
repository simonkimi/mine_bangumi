package config

import (
	"errors"
	"github.com/simonkimi/minebangumi/tools/dir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"path/filepath"
)

var AppConfig *AppConfigModel

var configPath = ""

func setDefault(key string, env string, value any) {
	if env != "" {
		_ = viper.BindEnv(key, env)
	}
	viper.SetDefault(key, value)
}

func setDefaultConfig() {
	setDefault("is_first_run", "", true)
	// Server
	setDefault("server.ipv4_host", "MBG_SERVER_IPV4_HOST", "0.0.0.0")
	setDefault("server.ipv4_port", "MBG_SERVER_IPV4_PORT", "7962")
	setDefault("server.ipv6_host", "MBG_SERVER_IPV6_HOST", "[::1]")
	setDefault("server.ipv6_port", "MBG_SERVER_IPV6_PORT", "7962")
	// User
	setDefault("user.username", "MBG_USERNAME", "admin")
	setDefault("user.password", "MBG_PASSWORD", "admin")

	// Path
	setDefault("path.workdir", "MBG_WORKDIR", dir.GetConfigDir())

	// Downloader
	setDefault("downloader.client", "MBG_DOWNLOADER", "")

	// qBittorrent
	setDefault("downloader.qBittorrent.host", "MBG_QBITTORRENT_HOST", "http://localhost:8080")
	setDefault("downloader.qBittorrent.username", "MBG_QBITTORRENT_USERNAME", "admin")
	setDefault("downloader.qBittorrent.password", "MBG_QBITTORRENT_PASSWORD", "adminadmin")

	// Aria2
	setDefault("downloader.aria2.host", "MBG_ARIA2_HOST", "http://localhost:6800/jsonrpc")
	setDefault("downloader.aria2.token", "MBG_ARIA2_TOKEN", "")

	// Proxy
	setDefault("proxy.enable", "MBG_PROXY_ENABLED", false)
	setDefault("proxy.scheme", "MBG_PROXY_SCHEME", "")
	setDefault("proxy.host", "MBG_PROXY_HOST", "")
	setDefault("proxy.port", "MBG_PROXY_PORT", "")
	setDefault("proxy.use_auth", "MBG_PROXY_USE_AUTH", false)
	setDefault("proxy.username", "MBG_PROXY_USERNAME", "")
	setDefault("proxy.password", "MBG_PROXY_PASSWORD", "")

	// Tmdb
	setDefault("tmdb.api_key", "MBG_TMDB_API_KEY", "")
}

func Setup() {
	configPath = filepath.Join(dir.GetConfigDir(), "config.toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("bangumi")

	viper.SetConfigName(dir.GetConfigDir())
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	setDefaultConfig()

	initConfig := false
	err := viper.ReadInConfig()
	if errors.As(err, &viper.ConfigFileNotFoundError{}) {
		logrus.Warn("Config file not found, use default values")
		initConfig = true
	}
	var config AppConfigModel
	if err := viper.Unmarshal(&config); err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal config")
	}
	if initConfig {
		if err := viper.WriteConfigAs(configPath); err != nil {
			logrus.WithError(err).Fatal("Failed to write config file")
		}
		logrus.Info("Config file created")
	}
	AppConfig = &config
}

func SaveConfig() {
	if err := viper.WriteConfigAs(configPath); err != nil {
		logrus.WithError(err).Fatal("Failed to write config file")
	}
	logrus.Info("Config file saved")
}
