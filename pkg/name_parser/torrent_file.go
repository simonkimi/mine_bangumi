package name_parser

type TorrentEpisodeFile struct {
	Path    string
	Group   string
	Title   string
	Season  int
	Episode int
	Ext     string
}

type TorrentSubtitleFile struct {
	Path     string
	Group    string
	Title    string
	Season   int
	Language string
	Episode  int
	Ext      string
}
