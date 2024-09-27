package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
)

// ConfigUser is the resolver for the configUser field.
func (r *mutationResolver) ConfigUser(_ context.Context, input api.UserConfigInput) (*api.ConfigResult, error) {
	conf := r.mgr.GetConfig()
	if input.Username != nil {
		conf.SetString(config.UserUsername, *input.Username)
	}
	if input.Password != nil {
		conf.SetString(config.UserPassword, *input.Password)
	}
	conf.Save()
	return getConfigResult(conf), nil
}

func getUserConfig(conf config.Config) *api.UserConfigResult {
	return &api.UserConfigResult{
		Username: conf.GetString(config.UserUsername),
	}
}

func getConfigResult(conf config.Config) *api.ConfigResult {
	return &api.ConfigResult{
		User: getUserConfig(conf),
	}
}
