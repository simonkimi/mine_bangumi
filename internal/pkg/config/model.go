package config

type AppConfigModel struct {
	IsFirstRun bool              `mapstructure:"is_first_run"`
	User       *UserConfig       `mapstructure:"user"`
	Path       *PathConfig       `mapstructure:"path"`
	Downloader *DownloaderConfig `mapstructure:"downloader"`
	MikanProxy *MikanProxyConfig `mapstructure:"mikan_proxy"`
}

type MikanProxyConfig struct {
	Enable   bool   `mapstructure:"enable"`
	Scheme   string `mapstructure:"scheme"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	UseAuth  bool   `mapstructure:"use_auth"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type PathConfig struct {
	Workdir string `mapstructure:"workdir"`
}

type DownloaderConfig struct {
	Client string `mapstructure:"client"`
}

type UserConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type QBittorrentConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Aria2Config struct {
	Host  string `mapstructure:"host"`
	Token string `mapstructure:"token"`
}
