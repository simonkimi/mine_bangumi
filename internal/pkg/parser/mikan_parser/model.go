package mikan_parser

type MikanBangumi struct {
	Title    string
	Episodes []*MikanEpisode
}

type MikanEpisode struct {
	Title       string
	Guid        string
	Torrent     string
	TorrentSize int
}
