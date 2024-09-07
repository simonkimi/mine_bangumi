package test

import (
	"github.com/joho/godotenv"
	"github.com/simonkimi/minebangumi/pkg/downloader"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(logrus.DebugLevel)
}

func TestQbDownloader(t *testing.T) {
	qb := downloader.NewQBittorrentDl("http://192.168.2.3:28080", os.Getenv("QB_USERNAME"), os.Getenv("QB_PASSWORD"))
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
