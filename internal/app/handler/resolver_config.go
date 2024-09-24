package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
)

// ConfigUser is the resolver for the configUser field.
func (r *mutationResolver) ConfigUser(_ context.Context, input api.UserConfigInput) (*api.ConfigResult, error) {
	if input.Username != nil {
		config.UserUsername.Set(*input.Username)
	}
	if input.Password != nil {
		config.UserPassword.Set(*input.Password)
	}
	return getConfigResult(), nil
}

func getUserConfig() *api.UserConfigResult {
	return &api.UserConfigResult{
		Username: config.UserUsername.Get(),
	}
}

func getConfigResult() *api.ConfigResult {
	return &api.ConfigResult{
		User: getUserConfig(),
	}
}
