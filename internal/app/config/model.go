package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type configItem[T any] struct {
	key          string
	env          string
	defaultValue T
}

func newConfigItem[T any](key string, env string, defaultValue T) *configItem[T] {
	return &configItem[T]{key: key, env: env, defaultValue: defaultValue}
}

func newConfigItemFunc[T any](key string, env string, defaultValueFunc func() T) *configItem[T] {
	return &configItem[T]{key: key, env: env, defaultValue: defaultValueFunc()}
}

func (c *configItem[T]) getString() string {
	return viper.GetString(c.key)
}

func (c *configItem[T]) register(v *viper.Viper) {
	if c.env != "" {
		err := v.BindEnv(c.key, c.env)
		if err != nil {
			logrus.Errorf("Config init bind env error: %s", err)
		}
	}

	v.SetDefault(c.key, c.defaultValue)
}
