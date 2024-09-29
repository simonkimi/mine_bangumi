package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
)

// ConfigUser is the resolver for the configUser field.
func (r *mutationResolver) ConfigUser(_ context.Context, input api.UserConfigInput) (*api.ConfigResult, error) {
	if err := api.Validate(
		api.V("username", input.Username, "omitempty,ascii,min=3,max=20"),
		api.V("password", input.Password, "omitempty,ascii,max=40,ascii"),
	); err != nil {
		return nil, err
	}

	conf := r.mgr.GetConfig()
	config.UpdateUser(conf, input.Username, input.Password)
	return getConfigResult(conf), nil
}

func getUserConfig(conf config.Config) *api.UserConfigResult {
	return &api.UserConfigResult{
		Username: conf.GetString(config.UserUsername),
		Token:    conf.GetString(config.UserApiToken),
	}
}

func getConfigResult(conf config.Config) *api.ConfigResult {
	return &api.ConfigResult{
		User: getUserConfig(conf),
	}
}
