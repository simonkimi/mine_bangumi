package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/tools/xstring"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initUserResolver() (config.Config, *Resolver) {
	mgr := new(service.MockManager)
	conf := config.NewMockConfig()
	mgr.EXPECT().GetConfig().Return(conf)
	return conf, newResolver(mgr)
}

func TestConfigUser(t *testing.T) {
	conf, r := initUserResolver()
	username := xstring.RandomString(5)
	password := xstring.RandomString(40)
	input := api.UserConfigInput{
		Username: &username,
		Password: &password,
	}
	result, err := r.Mutation().ConfigUser(context.Background(), input)
	assert.Nil(t, err)
	assert.Equal(t, username, result.User.Username)
	assert.Equal(t, username, conf.GetString(config.UserUsername))
	assert.Equal(t, password, conf.GetString(config.UserPassword))
	t.Logf("Token: %s", result.User.Token)
}
