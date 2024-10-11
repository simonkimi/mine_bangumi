package downloader

import (
	"github.com/simonkimi/minebangumi/pkg/testutil"
	"os"
	"testing"
)

func init() {
	testutil.LoadTestEnv()
}

func TestQbDownloader(t *testing.T) {
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
