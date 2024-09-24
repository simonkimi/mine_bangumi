package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
}

type configItem[T any] struct {
	key          string
	env          string
	defaultValue T
}

func newConfigItem[T any](key string, env string, defaultValue T) *configItem[T] {
	return &configItem[T]{key: key, env: env, defaultValue: defaultValue}
}

func (c *configItem[T]) getString() string {
	return viper.GetString(c.key)
}

func (c *configItem[T]) register() {
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

func (c *configItem[T]) Get() T {
	return viper.Get(c.key).(T)
}

func (c *configItem[T]) Set(value T) {
	viper.Set(c.key, value)
}
