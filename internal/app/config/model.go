package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type AppConfigModel struct {
	System      SystemConfig       `toml:"system"`
	User        UserConfig         `toml:"user"`
	Server      ServerConfig       `toml:"server"`
	Downloader  DownloaderConfig   `toml:"downloader"`
	ProxyConfig NetworkProxyConfig `toml:"proxy"`
	Tmdb        TmdbConfig         `toml:"tmdb"`
}

type NetworkProxyConfig struct {
	Enable   bool   `toml:"enable" env:"MBG_PROXY_ENABLED" default:"false"`
	Scheme   string `toml:"scheme" env:"MBG_PROXY_SCHEME" default:"http"`
	Host     string `toml:"host" env:"MBG_PROXY_HOST" default:"127.0.0.1"`
	Port     string `toml:"port" env:"MBG_PROXY_PORT" default:"7890"`
	UseAuth  bool   `toml:"use_auth" env:"MBG_PROXY_USE_AUTH"`
	Username string `toml:"username" env:"MBG_PROXY_USERNAME"`
	Password string `toml:"password" env:"MBG_PROXY_PASSWORD"`
}

type ServerConfig struct {
	Ipv4Host string `toml:"ipv4_host" default:"0.0.0.0" env:"MBG_SERVER_IPV4_HOST"`
	Ipv4Port int    `toml:"Ipv4_port" default:"7962" env:"MBG_SERVER_IPV4_PORT"`
	Ipv6Host string `toml:"ipv6_host" default:"[::1]" env:"MBG_SERVER_IPV6_HOST"`
	Ipv6Port int    `toml:"Ipv6_port" default:"7962" env:"MBG_SERVER_IPV6_PORT"`
}

type DownloaderConfig struct {
	Client      string            `toml:"client" env:"MBG_DOWNLOADER_CLIENT"`
	QBittorrent QBittorrentConfig `toml:"qBittorrent"`
	Aria2       Aria2Config       `toml:"aria2"`
}

type UserConfig struct {
	Username string `toml:"username" default:"admin" env:"MBG_USERNAME"`
	Password string `toml:"password" default:"admin" env:"MBG_PASSWORD"`
}

type QBittorrentConfig struct {
	Host     string `toml:"host" default:"http://127.0.0.1:8080" env:"MBG_QBITTORRENT_HOST"`
	Username string `toml:"username" env:"MBG_QBITTORRENT_USERNAME"`
	Password string `toml:"password" env:"MBG_QBITTORRENT_PASSWORD"`
}

type Aria2Config struct {
	Host  string `toml:"host" default:"http://localhost:6800/jsonrpc" env:"MBG_ARIA2_HOST"`
	Token string `toml:"token" env:"MBG_ARIA2_TOKEN" env:"MBG_ARIA2_TOKEN"`
}

type TmdbConfig struct {
	ApiKey string `toml:"api_key" env:"MBG_TMDB_API_KEY"`
}

type SystemConfig struct {
	SecretKey  string `toml:"secret"`
	IsFirstRun bool   `toml:"is_first_run"`
}

func initConfigStruct(model any) {
	v := reflect.ValueOf(model).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Struct {
			initConfigStruct(field.Addr().Interface())
			continue
		}

		sField := t.Field(i)
		isSet := false
		envTag := sField.Tag.Get("env")
		defaultTag := sField.Tag.Get("default")

		if envTag != "" {
			if value, exist := os.LookupEnv(envTag); exist {
				isSet = true
				switch field.Kind() {
				case reflect.String:
					field.SetString(value)
				case reflect.Bool:
					field.SetBool(guessBool(value))
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					intValue, err := strconv.Atoi(value)
					if err != nil {
						logrus.Warnf("Config init format error: %s", err)
						isSet = false
					} else {
						field.SetInt(int64(intValue))
					}
				default:
					logrus.Warnf("Config init unsupported type: %s", field.Kind())
					isSet = false
				}
			}
		}

		if !isSet && defaultTag != "" {
			switch field.Kind() {
			case reflect.String:
				field.SetString(defaultTag)
			case reflect.Bool:
				field.SetBool(guessBool(defaultTag))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intValue, err := strconv.Atoi(defaultTag)
				if err != nil {
					logrus.Warnf("Config init format error: %s", err)
				} else {
					field.SetInt(int64(intValue))
				}
			default:
				logrus.Warnf("Config init unsupported type: %s", field.Kind())
			}
		}
	}
}

func guessBool(value string) bool {
	value = strings.ToLower(strings.TrimSpace(value))
	return value == "true" || value == "1" || value == "yes" || value == "y" || value == "on" || value == "enable" || value == "enabled" || value == "ok"
}
