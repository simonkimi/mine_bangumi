package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service/user_service"
	"github.com/simonkimi/minebangumi/pkg/errno"
)

func Login(c *gin.Context) {
	var form api.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		_ = c.Error(errno.NewFormError(err))
		return
	}
	if form.Username != config.AppConfig.User.Username || form.Password != config.AppConfig.User.Password {
		_ = c.Error(errno.NewApiError(errno.ErrorUserPasswordWrong))
		return
	}
	token, err := user_service.GenerateJwt(form.Username)
	if err != nil {
		_ = c.Error(err)
		return
	}
	api.OkResponse(c, &api.TokenResponse{
		Token: token,
	})
}
