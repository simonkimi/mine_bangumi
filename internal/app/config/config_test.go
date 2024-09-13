package config

import (
	"github.com/simonkimi/minebangumi/pkg/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {

}

func TestSetup(t *testing.T) {
	tests.WorkOnTempDir(t, false)
	Setup()
	assert.Equal(t, AppConfig.User.Username, "admin")
	assert.Equal(t, AppConfig.User.Password, "admin")
}
