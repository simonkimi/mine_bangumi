package config

import (
	"github.com/simonkimi/minebangumi/pkg/hash"
	"github.com/spf13/viper"
)

var UserUsername = newConfigItem("user.username", "MBG_USER_USERNAME", "admin")
var UserPassword = newConfigItem("user.password", "MBG_USER_PASSWORD", "admin")
var UserApiToken = newConfigItemFunc("user.api_token", "MBG_USER_API_TOKEN", func() string {
	return hash.GenerateRandomKey(40)
})

var ServerHost = newConfigItem("server.host", "MBG_SERVER_HOST", "0.0.0.0")
var ServerPort = newConfigItem("server.port", "MBG_SERVER_PORT", 8080)

var DownloaderClient = newConfigItem("downloader.client", "MBG_DOWNLOADER_CLIENT", "")
var QBittorrentApi = newConfigItem("downloader.qbittorrent.api", "MBG_QB_API", "")
var QBittorrentUser = newConfigItem("downloader.qbittorrent.user", "MBG_QB_USER", "")
var QBittorrentPassword = newConfigItem("downloader.qbittorrent.password", "MBG_QB_PASSWORD", "")

var Aria2Api = newConfigItem("downloader.aria2.api", "MBG_ARIA2_API", "")
var Aria2Token = newConfigItem("downloader.aria2.token", "MBG_ARIA2_TOKEN", "")

var ProxyEnabled = newConfigItem("proxy.enabled", "MBG_PROXY_ENABLED", false)
var ProxyScheme = newConfigItem("proxy.scheme", "MBG_PROXY_SCHEME", "http")
var ProxyHost = newConfigItem("proxy.host", "MBG_PROXY_HOST", "")
var ProxyPort = newConfigItem("proxy.port", "MBG_PROXY_PORT", "")
var ProxyUsername = newConfigItem("proxy.username", "MBG_PROXY_USERNAME", "")
var ProxyPassword = newConfigItem("proxy.password", "MBG_PROXY_PASSWORD", "")

var TmdbApiKey = newConfigItem("tmdb.api_key", "MBG_TMDB_KEY", "")

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
}
