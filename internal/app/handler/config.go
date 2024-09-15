package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/domain"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"net/http"
)

func System(c *gin.Context) {
	user := middleware.GetClaims(c)
	api.OkResponse(c, &api.SystemInfo{
		Version:    domain.Version,
		IsFirstRun: config.AppConfig.System.IsFirstRun,
		IsLogin:    user != nil,
	})
}

func InitUser(c *gin.Context) {
	if !config.AppConfig.System.IsFirstRun {
		_ = c.Error(errno.NewApiError(http.StatusForbidden))
		return
	}
	var form api.InitUserForm
	if err := c.ShouldBindJSON(&form); err != nil {
		_ = c.Error(errno.NewFormError(err))
	}
	config.UpdateUser(form.Username, form.Password)

	token, err := service.GenerateUserJwt(form.Username)
	if err != nil {
		_ = c.Error(errno.NewApiErrorWithCause(http.StatusInternalServerError, err))
		return
	}
	api.OkResponse(c, &api.TokenResponse{
		Token: token,
	})
}
