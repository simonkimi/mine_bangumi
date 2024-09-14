package config

import (
	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var AppConfig AppConfigModel

var configPath = ""

func Setup() {
	wd, err := os.Getwd()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get working directory")
	}
	configPath = filepath.Join(wd, "config.toml")

	viper.SetConfigName("config")
	viper.AddConfigPath(wd)
	viper.SetConfigType("toml")

	initConfig := false
	var config AppConfigModel
	setViperDefault(&config, []string{})

	err = viper.ReadInConfig()
	var configFileNotFound viper.ConfigFileNotFoundError
	if errors.As(err, &configFileNotFound) {
		logrus.Warn("Config file not found, use default values")
		initConfig = true
	}

	if err := viper.Unmarshal(&config); err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal config")
	}
	if initConfig {
		if err := viper.WriteConfigAs(configPath); err != nil {
			logrus.WithError(err).Fatal("Failed to write config file")
		}
		logrus.Info("Config file created")
	}
	AppConfig = config
}

func saveConfig() {
	if err := viper.WriteConfigAs(configPath); err != nil {
		logrus.WithError(err).Fatal("Failed to write config file")
	}
}
