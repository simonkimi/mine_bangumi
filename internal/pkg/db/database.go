package db

import (
	"github.com/simonkimi/minebangumi/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Setup() {
	var conn gorm.Dialector
	switch config.AppConfig.Database.Backends {
	case "sqlite":
		conn = sqlite.Open(config.AppConfig.Sqlite.Path)
	default:
		logrus.Fatalf("Unsupported database backends: %s", config.AppConfig.Database.Backends)
	}
	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to database")
	}
	Db = db
}
