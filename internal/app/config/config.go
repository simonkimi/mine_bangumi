package config

import (
	"github.com/cockroachdb/errors"
	"github.com/simonkimi/minebangumi/pkg/secret"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var AppConfig AppConfigModel

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

	initConfig := false
	var config AppConfigModel
	setViperDefault(&config, []string{})
	if defaultInit != nil {
		for k, v := range defaultInit {
			viper.SetDefault(k, v)
		}
	}

	err = viper.ReadInConfig()
	var configFileNotFound viper.ConfigFileNotFoundError
	if errors.As(err, &configFileNotFound) {
		logrus.Warn("Config file not found, use default values")
		initConfig = true
		config.System.SecretKey = secret.GenerateRandomKey(50)
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

func InitUser(username string, password string) {
	viper.Set("user.username", username)
	viper.Set("user.password", password)
	viper.Set("user.init_user", false)
	AppConfig.User.Username = username
	AppConfig.User.Password = password
	AppConfig.User.IsInitUser = true
	saveConfig()
}

func saveConfig() {
	if err := viper.WriteConfigAs(configPath); err != nil {
		logrus.WithError(err).Fatal("Failed to write config file")
	}
}
