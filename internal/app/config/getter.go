package config

func GetUserUsername() string {
	return config.keys.userUsername.getString()
}

func GetUserPassword() string {
	return config.keys.userPassword.getString()
}

func GetServerIpv4Host() string {
	return config.keys.serverIpv4Host.getString()
}

func GetServerIpv4Port() string {
	return config.keys.serverIpv4Port.getString()
}

func GetServerIpv6Host() string {
	return config.keys.serverIpv6Host.getString()
}

func GetServerIpv6Port() string {
	return config.keys.serverIpv6Port.getString()
}

func GetDownloaderClient() string {
	return config.keys.downloaderClient.getString()
}

func GetQBittorrentApi() string {
	return config.keys.qBittorrentApi.getString()
}

func GetQBittorrentUser() string {
	return config.keys.qBittorrentUser.getString()
}

func GetQBittorrentPassword() string {
	return config.keys.qBittorrentPassword.getString()
}

func GetAria2Api() string {
	return config.keys.aria2Api.getString()
}

func GetAria2Token() string {
	return config.keys.aria2Token.getString()
}

func GetTmdbApiKey() string {
	return config.keys.tmdbApiKey.getString()
}
