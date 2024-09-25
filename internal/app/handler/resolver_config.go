package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/manager"
)

// ConfigUser is the resolver for the configUser field.
func (r *mutationResolver) ConfigUser(_ context.Context, input api.UserConfigInput) (*api.ConfigResult, error) {
	mgr := manager.GetInstance()
	if input.Username != nil {
		mgr.Config.SetString(config.UserUsername, *input.Username)
	}
	if input.Password != nil {
		mgr.Config.SetString(config.UserPassword, *input.Password)
	}
	mgr.Config.Save()
	return getConfigResult(), nil
}

func getUserConfig() *api.UserConfigResult {
	mgr := manager.GetInstance()
	return &api.UserConfigResult{
		Username: mgr.Config.GetString(config.UserUsername),
	}
}

func getConfigResult() *api.ConfigResult {
	return &api.ConfigResult{
		User: getUserConfig(),
	}
}
