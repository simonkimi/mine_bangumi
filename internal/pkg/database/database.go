package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Setup() {
	conn := sqlite.Open("app.sqlite3")
	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to database")
	}
	Db = db
}
