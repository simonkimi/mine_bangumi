package mikan

type Bangumi struct {
	Title    string
	Episodes []*Episode
}

type Episode struct {
	Title       string
	Guid        string
	Link        string
	Torrent     string
	Description string
	TorrentSize int
}
