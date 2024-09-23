package config

import (
	"github.com/simonkimi/minebangumi/pkg/logger"
	"github.com/simonkimi/minebangumi/pkg/tests"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	tests.MainOnTempDir()
	logger.Setup()
	tests.LoadTestEnv()
	m.Run()
}

func TestSetup(t *testing.T) {
	Setup()
	assert.Equal(t, appConfig.User.Username, "admin")
	assert.Equal(t, appConfig.User.Password, "admin")
	viper.Debug()
}
