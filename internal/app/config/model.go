package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"reflect"
	"strconv"
	"strings"
)

type AppConfigModel struct {
	System      SystemConfig       `mapstructure:"system"`
	User        UserConfig         `mapstructure:"user"`
	Server      ServerConfig       `mapstructure:"server"`
	Downloader  DownloaderConfig   `mapstructure:"downloader"`
	ProxyConfig NetworkProxyConfig `mapstructure:"proxy"`
	Tmdb        TmdbConfig         `mapstructure:"tmdb"`
}

type NetworkProxyConfig struct {
	Enable   bool   `mapstructure:"enable" env:"MBG_PROXY_ENABLED" default:"false"`
	Scheme   string `mapstructure:"scheme" env:"MBG_PROXY_SCHEME" default:"http"`
	Host     string `mapstructure:"host" env:"MBG_PROXY_HOST" default:"127.0.0.1"`
	Port     string `mapstructure:"port" env:"MBG_PROXY_PORT" default:"7890"`
	UseAuth  bool   `mapstructure:"use_auth" env:"MBG_PROXY_USE_AUTH" default:"false"`
	Username string `mapstructure:"username" env:"MBG_PROXY_USERNAME" default:""`
	Password string `mapstructure:"password" env:"MBG_PROXY_PASSWORD" default:""`
}

type ServerConfig struct {
	Ipv4Host string `mapstructure:"ipv4_host" env:"MBG_SERVER_IPV4_HOST" default:"0.0.0.0"`
	Ipv4Port int    `mapstructure:"Ipv4_port" env:"MBG_SERVER_IPV4_PORT" default:"7962"`
	Ipv6Host string `mapstructure:"ipv6_host" env:"MBG_SERVER_IPV6_HOST" default:""`
	Ipv6Port int    `mapstructure:"Ipv6_port" env:"MBG_SERVER_IPV6_PORT" default:""`
}

type DownloaderConfig struct {
	Client      string            `mapstructure:"client" env:"MBG_DOWNLOADER_CLIENT" default:""`
	QBittorrent QBittorrentConfig `mapstructure:"qBittorrent" default:""`
	Aria2       Aria2Config       `mapstructure:"aria2" default:""`
}

type UserConfig struct {
	Username string `mapstructure:"username" env:"MBG_USERNAME" default:"admin"`
	Password string `mapstructure:"password" env:"MBG_PASSWORD" default:"admin"`
}

type QBittorrentConfig struct {
	Host     string `mapstructure:"host" env:"MBG_QBITTORRENT_HOST" default:"http://127.0.0.1:8080"`
	Username string `mapstructure:"username" env:"MBG_QBITTORRENT_USERNAME" default:""`
	Password string `mapstructure:"password" env:"MBG_QBITTORRENT_PASSWORD" default:""`
}

type Aria2Config struct {
	Host  string `mapstructure:"host"  env:"MBG_ARIA2_HOST" default:"http://localhost:6800/jsonrpc"`
	Token string `mapstructure:"token" env:"MBG_ARIA2_TOKEN" default:""`
}

type TmdbConfig struct {
	ApiKey string `mapstructure:"api_key" env:"MBG_TMDB_API_KEY" default:""`
}

type SystemConfig struct {
	SecretKey  string `mapstructure:"secret" default:""`
	IsFirstRun bool   `mapstructure:"is_first_run" default:"true"`
}

func setViperDefault(model any, path []string) {
	v := reflect.ValueOf(model).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		vField := v.Field(i)
		sField := t.Field(i)

		mapper := sField.Tag.Get("mapstructure")
		elePath := strings.Join(append(path, mapper), ".")

		if vField.Kind() == reflect.Struct {
			nextPath := append(path, mapper)
			setViperDefault(vField.Addr().Interface(), nextPath)
			continue
		}
		envTag := sField.Tag.Get("env")
		if envTag != "" {
			err := viper.BindEnv(elePath, envTag)
			if err != nil {
				logrus.Errorf("Config init bind env error: %s", err)
			}
		}

		defaultTag := sField.Tag.Get("default")
		if defaultTag != "" {
			switch vField.Kind() {
			case reflect.String:
				viper.SetDefault(elePath, defaultTag)
			case reflect.Bool:
				viper.SetDefault(elePath, guessBool(defaultTag))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value, err := strconv.Atoi(defaultTag)
				if err != nil {
					logrus.Errorf("Config init format error: %s", err)
				}
				viper.SetDefault(elePath, value)
			case reflect.Float64, reflect.Float32:
				value, err := strconv.ParseFloat(defaultTag, 64)
				if err != nil {
					logrus.Errorf("Config init format error: %s", err)
				}
				viper.SetDefault(elePath, value)
			default:
				logrus.Errorf("Config init unsupported type: %s", vField.Kind())
			}
		}

	}
}

func guessBool(value string) bool {
	value = strings.ToLower(strings.TrimSpace(value))
	return value == "true" || value == "1" || value == "yes" || value == "y" || value == "on" || value == "enable" || value == "enabled" || value == "ok"
}
