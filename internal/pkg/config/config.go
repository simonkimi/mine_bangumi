package config

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var AppConfig *AppConfigModel

const configName = "config.toml"

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
	viper.AutomaticEnv()
	viper.SetEnvPrefix("bangumi")
	setDefault("database.backends", "DATABASE_BACKENDS", "sqlite3")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

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
		if err := viper.WriteConfigAs(configName); err != nil {
			logrus.WithError(err).Fatal("Failed to write config file")
		}
		logrus.Info("Config file created")
	}
	AppConfig = &config
}

func SaveConfig() {
	if err := viper.WriteConfigAs(configName); err != nil {
		logrus.WithError(err).Fatal("Failed to write config file")
	}
	logrus.Info("Config file saved")
}
