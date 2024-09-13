package downloader

type IDownloader interface {
	Login() error
	RecordClientInfo() (string, error)
}
