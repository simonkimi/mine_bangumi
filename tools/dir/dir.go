package dir

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const appDomain = "ink.z31.minebangumi"

func Setup() {
	createDirIfNotExist(GetConfigDir())
	createDirIfNotExist(GetTempDir())
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			logrus.WithError(err).Fatalf("Failed to create dir %s", dir)
		}
	}
}

func GetConfigDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get user config dir")
	}

	return filepath.Join(dir, appDomain)
}

func GetTempDir() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get user cache dir")
	}

	return filepath.Join(dir, appDomain)
}
