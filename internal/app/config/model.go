package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Item[T any] struct {
	key          string
	env          string
	defaultValue T
}

func newConfig[T any](key string, defaultValue T) *Item[T] {
	return &Item[T]{key: key, env: "", defaultValue: defaultValue}
}

func newConfigEnv[T any](key string, env string, defaultValue T) *Item[T] {
	return &Item[T]{key: key, env: env, defaultValue: defaultValue}
}

func newConfigItemFunc[T any](key string, env string, defaultValueFunc func() T) *Item[T] {
	return &Item[T]{key: key, env: env, defaultValue: defaultValueFunc()}
}

func (c *Item[T]) getString() string {
	return viper.GetString(c.key)
}

func (c *Item[T]) register(v *viper.Viper) {
	if c.env != "" {
		err := v.BindEnv(c.key, c.env)
		if err != nil {
			logrus.Errorf("ConfigService init bind env error: %s", err)
		}
	}

	v.SetDefault(c.key, c.defaultValue)
}
