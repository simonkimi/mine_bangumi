package db

import (
	"github.com/simonkimi/minebangumi/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
)

var Db *gorm.DB

func Setup() {
	conn := sqlite.Open(filepath.Join(config.AppConfig.Path.Workdir, "app.sqlite3"))
	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to database")
	}
	Db = db
}
