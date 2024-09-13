package config

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
	"os"
)

var AppConfig AppConfigModel

var configPath = ""

func Setup() {
	configPath = "config.toml"
	var config AppConfigModel

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logrus.Warn("Config file not found, use default values")
		initConfigStruct(&config)
		AppConfig = config
		SaveConfig()
		return
	}

	fileData, err := os.ReadFile(configPath)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to read config file")
	}
	err = toml.Unmarshal(fileData, &config)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal config")
	}

	AppConfig = config
}

func SaveConfig() {
	config, err := toml.Marshal(AppConfig)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to marshal config")
	}
	err = os.WriteFile(configPath, config, 0644)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to write config file")
	}
}
