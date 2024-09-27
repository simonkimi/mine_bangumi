package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

//go:generate mockery --name Config
type Config interface {
	GetString(key *configItem[string]) string
	GetInt(key *configItem[int]) int
	GetBool(key *configItem[bool]) bool
	GetFloat64(key *configItem[float64]) float64
	SetString(key *configItem[string], value string)
	SetInt(key *configItem[int], value int)
	SetBool(key *configItem[bool], value bool)
	Save()
}

type Service struct {
	viper      *viper.Viper
	configPath string
	mu         sync.Locker
}

func NewConfig() (Config, error) {
	v := viper.New()
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	v.SetConfigFile(path)
	registerKey(v)
	config := &Service{
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
		fmt.Println("ConfigService file not found, use default values")
	}

	return path, nil
}

func (c *Service) GetString(key *configItem[string]) string {
	return c.viper.GetString(key.key)
}

func (c *Service) GetInt(key *configItem[int]) int {
	return c.viper.GetInt(key.key)
}

func (c *Service) GetBool(key *configItem[bool]) bool {
	return c.viper.GetBool(key.key)
}

func (c *Service) GetFloat64(key *configItem[float64]) float64 {
	return c.viper.GetFloat64(key.key)
}

func (c *Service) SetString(key *configItem[string], value string) {
	c.viper.Set(key.key, value)
}

func (c *Service) SetInt(key *configItem[int], value int) {
	c.viper.Set(key.key, value)
}

func (c *Service) SetBool(key *configItem[bool], value bool) {
	c.viper.Set(key.key, value)
}

func (c *Service) Save() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if err := viper.WriteConfigAs(c.configPath); err != nil {
		fmt.Printf("Failed to write config file, %+v", err)
	}
}
