package config

type AppConfigModel struct {
	IsFirstRun bool              `mapstructure:"is_first_run"`
	User       *UserConfig       `mapstructure:"user"`
	Server     *ServerConfig     `mapstructure:"server"`
	Path       *PathConfig       `mapstructure:"path"`
	Downloader *DownloaderConfig `mapstructure:"downloader"`
	MikanProxy *MikanProxyConfig `mapstructure:"mikan_proxy"`
}

type ServerConfig struct {
	Ipv4Host string `mapstructure:"ipv4_host"`
	Ipv4Port string `mapstructure:"Ipv4_port"`
	Ipv6Host string `mapstructure:"ipv6_host"`
	Ipv6Port string `mapstructure:"Ipv6_port"`
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
