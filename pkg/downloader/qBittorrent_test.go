package downloader

import (
	"os"
	"testing"
)

func init() {
	testutil.LoadTestEnv()
}

func TestQbDownloader(t *testing.T) {
	testutil.WorkOnTempDir(t, true)

	host := os.Getenv("MBG_QBITTORRENT_HOST")
	username := os.Getenv("MBG_QBITTORRENT_USERNAME")
	password := os.Getenv("MBG_QBITTORRENT_PASSWORD")

	qb := NewQBittorrentDl(host, username, password)
	err := qb.Login()
	if err != nil {
		t.Fatal(err)
	}

	version, err := qb.RecordClientInfo()
	if err != nil {
		return
	}
	t.Log(version)
}
