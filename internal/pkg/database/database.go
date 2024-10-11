package database

import (
	"database/sql"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Db            *gorm.DB
	rawDb         *sql.DB
	schemaVersion uint
}

func NewDb(path string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}
	rawDb, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get raw database")
	}

	m, err := newMigrator(rawDb)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create migrator")
	}

	database := &Database{
		Db:            db,
		rawDb:         rawDb,
		schemaVersion: m.getSchemaVersion(),
	}

	if m.isNewDb() {
		if err := m.m.Up(); err != nil {
			return nil, errors.Wrap(err, "failed to run migrations")
		}
		database.schemaVersion = m.getSchemaVersion()
		err := database.Optimise()
		if err != nil {
			return nil, err
		}
	} else {
		currentVersion := m.getSchemaVersion()
		if currentVersion > appSchemaVersion {
			return nil, errors.Errorf("database schema version is newer than expected: %d > %d", currentVersion, appSchemaVersion)
		}
	}

	return database, nil
}

func (d *Database) GetSchemaVersion() uint {
	return d.schemaVersion
}

func (d *Database) GetAppSchemaVersion() uint {
	return appSchemaVersion
}

func (d *Database) NeedMigrate() bool {
	return d.schemaVersion < appSchemaVersion
}

func (d *Database) Optimise() error {
	_, err := d.rawDb.Exec("ANALYZE")
	if err != nil {
		return errors.Wrap(err, "failed to analyze database")
	}
	_, err = d.rawDb.Exec("VACUUM")
	if err != nil {
		return errors.Wrap(err, "failed to vacuum database")
	}
	return nil
}

func (d *Database) Close() error {
	return d.rawDb.Close()
}
