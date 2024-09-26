package database

import (
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDb() (*gorm.DB, error) {
	conn := sqlite.Open("app.sqlite3")
	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}
	return db, nil
}
