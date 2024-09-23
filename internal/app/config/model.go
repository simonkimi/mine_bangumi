package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const DownloaderTypeAria2 = "aria2"
const DownloaderTypeQBittorrent = "qbittorrent"

type appConfig struct {
	IsNewSystem bool
	keys        *configKeys
}

type configItem struct {
	key          string
	env          string
	defaultValue any
}

func newConfigItem(key string, env string, defaultValue any) *configItem {
	return &configItem{key: key, env: env, defaultValue: defaultValue}
}

func (c *configItem) getString() string {
	return viper.GetString(c.key)
}

func (c *configItem) register() {
	if c.env != "" {
		err := viper.BindEnv(c.key, c.env)
		if err != nil {
			logrus.Errorf("Config init bind env error: %s", err)
		}
	}
	if c.defaultValue != nil {
		viper.SetDefault(c.key, c.defaultValue)
	}
}

func (c *configItem) getInt() int {
	return viper.GetInt(c.key)
}

func (c *configItem) getBool() bool {
	return viper.GetBool(c.key)
}

func (c *configItem) getFloat64() float64 {
	return viper.GetFloat64(c.key)
}

func (c *configItem) setValue(value any) {
	viper.Set(c.key, value)
}
