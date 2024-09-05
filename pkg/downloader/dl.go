package downloader

type DownloadOption struct {
	DownloadPath string
}

type IDownloader interface {
	StartTorrent(filepath string, option *DownloadOption) (id string, err error)
	StartMagnet(magnet string, option *DownloadOption) (id string, err error)
}
