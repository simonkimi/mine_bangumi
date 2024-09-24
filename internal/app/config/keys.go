package config

var UserUsername = newConfigItem("user.username", "MBG_USER_USERNAME", "admin")
var UserPassword = newConfigItem("user.password", "MBG_USER_PASSWORD", "admin")

var ServerIpv4Host = newConfigItem("server.host", "MBG_SERVER_HOST", "0.0.0.0")
var ServerIpv4Port = newConfigItem("server.port", "MBG_SERVER_PORT", 8080)

var DownloaderClient = newConfigItem("downloader.client", "MBG_DOWNLOADER_CLIENT", "")

var QBittorrentApi = newConfigItem("downloader.qbittorrent.api", "MBG_QB_API", "")
var QBittorrentUser = newConfigItem("downloader.qbittorrent.user", "MBG_QB_USER", "")
var QBittorrentPassword = newConfigItem("downloader.qbittorrent.password", "MBG_QB_PASSWORD", "")

var Aria2Api = newConfigItem("downloader.aria2.api", "MBG_ARIA2_API", "")
var Aria2Token = newConfigItem("downloader.aria2.token", "MBG_ARIA2_TOKEN", "")

var TmdbApiKey = newConfigItem("tmdb.api_key", "MBG_TMDB_KEY", "")
