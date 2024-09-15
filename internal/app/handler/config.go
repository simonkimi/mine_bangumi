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

// System godoc
// @Summary Get system information
// @Description Get the current system version, first run status, and login status
// @Tags system
// @Accept json
// @Produce json
// @Success 200 {object} api.SystemInfo
// @Router /api/v1/config/system [get]
func System(c *gin.Context) {
	user := middleware.GetClaims(c)
	api.OkResponse(c, &api.SystemInfo{
		Version:    domain.Version,
		IsFirstRun: config.AppConfig.System.IsFirstRun,
		IsLogin:    user != nil,
	})
}

// InitUser godoc
// @Summary Initialize the first user
// @Description Initialize the first user with a username and password
// @Tags user
// @Accept json
// @Produce json
// @Param InitUserForm body api.InitUserForm true "User initialization form"
// @Success 200 {object} api.TokenResponse
// @Failure 400 {object} errno.ApiError "Invalid form data"
// @Failure 403 {object} errno.ApiError "Forbidden"
// @Failure 500 {object} errno.ApiError "Internal server error"
// @Router /api/v1/config/init_user [post]
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
