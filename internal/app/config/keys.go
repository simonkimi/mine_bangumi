package config

import (
	"github.com/simonkimi/minebangumi/pkg/hash"
	"github.com/spf13/viper"
)

const appConfigVersion = 1

var SystemConfigVersion = newConfig("system.version", 0)
var SystemInit = newConfig("system.init", false)

var UserUsername = newConfigEnv("user.username", "MBG_USER_USERNAME", "admin")
var UserPassword = newConfigEnv("user.password", "MBG_USER_PASSWORD", "admin")
var UserApiToken = newConfigItemFunc("user.api_token", "MBG_USER_API_TOKEN", func() string {
	return hash.GenerateRandomKey(40)
})

var ServerHost = newConfigEnv("server.host", "MBG_SERVER_HOST", "0.0.0.0")
var ServerPort = newConfigEnv("server.port", "MBG_SERVER_PORT", 7962)

var DownloaderClient = newConfigEnv("downloader.client", "MBG_DOWNLOADER_CLIENT", "")
var QBittorrentApi = newConfigEnv("downloader.qbittorrent.api", "MBG_QB_API", "")
var QBittorrentUser = newConfigEnv("downloader.qbittorrent.user", "MBG_QB_USER", "")
var QBittorrentPassword = newConfigEnv("downloader.qbittorrent.password", "MBG_QB_PASSWORD", "")

var Aria2Api = newConfigEnv("downloader.aria2.api", "MBG_ARIA2_API", "")
var Aria2Token = newConfigEnv("downloader.aria2.token", "MBG_ARIA2_TOKEN", "")

var ProxyEnabled = newConfigEnv("proxy.enabled", "MBG_PROXY_ENABLED", false)
var ProxyScheme = newConfigEnv("proxy.scheme", "MBG_PROXY_SCHEME", "http")
var ProxyHost = newConfigEnv("proxy.host", "MBG_PROXY_HOST", "")
var ProxyPort = newConfigEnv("proxy.port", "MBG_PROXY_PORT", "")
var ProxyUsername = newConfigEnv("proxy.username", "MBG_PROXY_USERNAME", "")
var ProxyPassword = newConfigEnv("proxy.password", "MBG_PROXY_PASSWORD", "")

var TmdbApiKey = newConfigEnv("tmdb.api_key", "MBG_TMDB_KEY", "")

func registerKey(v *viper.Viper) {
	UserUsername.register(v)
	UserPassword.register(v)
	ServerHost.register(v)
	ServerPort.register(v)
	DownloaderClient.register(v)
	QBittorrentApi.register(v)
	QBittorrentUser.register(v)
	QBittorrentPassword.register(v)
	Aria2Api.register(v)
	Aria2Token.register(v)
	ProxyEnabled.register(v)
	ProxyScheme.register(v)
	ProxyHost.register(v)
	ProxyPort.register(v)
	ProxyUsername.register(v)
	ProxyPassword.register(v)
	TmdbApiKey.register(v)
	SystemInit.register(v)
	SystemConfigVersion.register(v)
}
