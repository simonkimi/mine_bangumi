package database

import (
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/pkg/errors"
)

const appSchemaVersion = 1

//go:embed migrations/*.sql
var migrationFiles embed.FS

type migrator struct {
	m *migrate.Migrate
}

func newMigrator(db *sql.DB) (*migrator, error) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create sqlite3 driver")
	}

	migrations, err := iofs.New(migrationFiles, "migrations")

	m, err := migrate.NewWithInstance("iofs", migrations, "sqlite3", driver)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create migrate instance")
	}
	return &migrator{m}, nil
}

func (m *migrator) getSchemaVersion() uint {
	version, _, _ := m.m.Version()
	return version
}

func (m *migrator) needsMigration() bool {
	return m.getSchemaVersion() != appSchemaVersion
}

func (m *migrator) isNewDb() bool {
	return m.getSchemaVersion() == 0
}

func (m *migrator) Migrate() error {
	err := m.m.Up()
	if err != nil {
		return err
	}
	return nil
}

func (m *migrator) Close() {
	_, _ = m.m.Close()
}
