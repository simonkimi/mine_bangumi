package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	viper      *viper.Viper
	configPath string
	mu         sync.Locker
}

func NewConfig() (*Config, error) {
	v := viper.New()
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	v.SetConfigFile(path)
	registerKey(v)
	config := &Config{
		viper:      v,
		configPath: path,
	}
	return config, nil
}

func getConfigPath() (string, error) {
	var path string
	if os.Getenv("MBG_CONFIG_PATH") != "" {
		path = os.Getenv("MBG_CONFIG_PATH")
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = filepath.Join(wd, "config.toml")
	}
	err := viper.ReadInConfig()
	var configFileNotFound viper.ConfigFileNotFoundError
	if errors.As(err, &configFileNotFound) {
		fmt.Println("Config file not found, use default values")
	}

	return path, nil
}

func (c *Config) GetString(key *configItem[string]) string {
	return c.viper.GetString(key.key)
}

func (c *Config) GetInt(key *configItem[int]) int {
	return c.viper.GetInt(key.key)
}

func (c *Config) GetBool(key *configItem[bool]) bool {
	return c.viper.GetBool(key.key)
}

func (c *Config) GetFloat64(key *configItem[float64]) float64 {
	return c.viper.GetFloat64(key.key)
}

func (c *Config) SetString(key *configItem[string], value string) {
	c.viper.Set(key.key, value)
}

func (c *Config) SetInt(key *configItem[int], value int) {
	c.viper.Set(key.key, value)
}

func (c *Config) SetBool(key *configItem[bool], value bool) {
	c.viper.Set(key.key, value)
}

func (c *Config) Save() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if err := viper.WriteConfigAs(c.configPath); err != nil {
		fmt.Printf("Failed to write config file, %+v", err)
	}
}
