package database

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewDb(t *testing.T) {
	dbName := "TestNewDb.sqlite"
	_ = os.Remove(dbName)
	t.Cleanup(func() {
		_ = os.Remove(dbName)
	})
	db, err := NewDb(dbName)
	defer db.Close()
	require.Nil(t, err)
	require.NotNil(t, db)

	require.Equal(t, uint(appSchemaVersion), db.GetSchemaVersion())
	require.True(t, db.Db.Migrator().HasTable("subscriptions"))
	require.True(t, db.Db.Migrator().HasTable("subscription_files"))
	require.True(t, db.Db.Migrator().HasTable("bangumis"))
	require.True(t, db.Db.Migrator().HasTable("bangumi_episodes"))
}
