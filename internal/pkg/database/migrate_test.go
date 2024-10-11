package database

import (
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestMigrator(t *testing.T) {
	testDb := "migrate_test.sqlite"
	t.Cleanup(func() {
		_ = os.Remove(testDb)
	})
	_ = os.Remove(testDb)
	sql, err := gorm.Open(sqlite.Open(testDb))
	require.Nil(t, err)
	db, err := sql.DB()
	defer db.Close()
	require.Nil(t, err)
	m, err := newMigrator(db)
	require.Nil(t, err)
	defer m.Close()

	require.True(t, m.needsMigration())
	require.True(t, m.isNewDb())

	err = m.Migrate()
	require.Nil(t, err)

	require.True(t, sql.Migrator().HasTable("subscriptions"))
	require.True(t, sql.Migrator().HasTable("subscription_files"))
	require.True(t, sql.Migrator().HasTable("bangumis"))
	require.True(t, sql.Migrator().HasTable("bangumi_episodes"))
}
