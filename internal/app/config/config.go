package config

func Initialize() {

}

//
//var configPath = ""
//
//func Setup() {
//	wd, err := os.Getwd()
//	if err != nil {
//		logrus.WithError(err).Fatal("Failed to get working directory")
//	}
//	configPath = filepath.Join(wd, "config.toml")
//
//	viper.SetConfigName("config")
//	viper.AddConfigPath(wd)
//	viper.SetConfigType("toml")
//
//	configs := []*configItem{
//		UserUsername,
//		UserPassword,
//		ServerIpv4Host,
//		ServerIpv4Port,
//		DownloaderClient,
//		QBittorrentApi,
//		QBittorrentUser,
//		QBittorrentPassword,
//		Aria2Api,
//		Aria2Token,
//		TmdbApiKey,
//	}
//
//	for _, c := range configs {
//		c.register()
//	}
//
//	err = viper.ReadInConfig()
//	var configFileNotFound viper.ConfigFileNotFoundError
//	if errors.As(err, &configFileNotFound) {
//		logrus.Warn("Config file not found, use default values")
//	}
//}
//
//func SaveConfig() {
//	if err := viper.WriteConfigAs(configPath); err != nil {
//		logrus.WithError(err).Fatal("Failed to write config file")
//	}
//}
