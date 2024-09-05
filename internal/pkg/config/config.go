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

type AppConfigModel struct {
	Database *DatabaseConfig `mapstructure:"database"`
	Sqlite   *SqliteConfig   `mapstructure:"sqlite"`
}

type DatabaseConfig struct {
	Backends string `mapstructure:"backends"`
}

type SqliteConfig struct {
	Path string `mapstructure:"path"`
}

func setDefault(key string, env string, value any) {
	_ = viper.BindEnv(key, env)
	viper.SetDefault(key, value)
}

func Setup() {
	configPath = filepath.Join(dir.GetConfigDir(), "config.toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("bangumi")

	viper.SetConfigName(dir.GetConfigDir())
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	setDefault("database.backends", "DATABASE_BACKENDS", "sqlite3")
	setDefault("sqlite.path", "SQLITE_PATH", filepath.Join(dir.GetConfigDir(), "bangumi.db"))

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
