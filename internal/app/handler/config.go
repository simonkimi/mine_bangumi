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

// GetSystem godoc
// @Summary Get system information
// @Description Get the current system version, first run status, and login status
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} api.SystemInfo
// @Router /api/v1/config/system [get]
func GetSystem(c *gin.Context) {
	user := middleware.GetClaims(c)
	api.OkResponse(c, &api.SystemInfo{
		Version:    domain.Version,
		IsInitUser: config.AppConfig.User.IsInitUser,
		IsLogin:    user != nil,
	})
}

// PostInitUser godoc
// @Summary Initialize the first user
// @Description Initialize the first user with a username and password
// @Tags config
// @Accept json
// @Produce json
// @Param InitUserForm body api.InitUserForm true "User initialization form"
// @Success 200 {object} api.TokenResponse
// @Failure 400 {object} errno.ApiError "Invalid form data"
// @Failure 403 {object} errno.ApiError "Forbidden"
// @Failure 500 {object} errno.ApiError "Internal server error"
// @Router /api/v1/config/init_user [post]
func PostInitUser(c *gin.Context) {
	if config.AppConfig.User.IsInitUser {
		_ = c.Error(errno.NewApiError(http.StatusForbidden))
		return
	}
	var form api.InitUserForm
	if err := c.ShouldBindJSON(&form); err != nil {
		_ = c.Error(errno.NewFormError(err))
	}
	config.InitUser(form.Username, form.Password)

	token, err := service.GenerateUserJwt(form.Username)
	if err != nil {
		_ = c.Error(errno.NewApiErrorWithCause(http.StatusInternalServerError, err))
		return
	}
	api.OkResponse(c, &api.TokenResponse{
		Token: token,
	})
}

// GetDownloader godoc
// @Summary Get the downloader configuration
// @Description Get the downloader configuration, including the type, API address, and token
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} api.DownloaderForm
// @Router /api/v1/config/downloader [get]
func GetDownloader(c *gin.Context) {
	form := &api.DownloaderForm{}
	form.Type = config.AppConfig.Downloader.Client
	switch form.Type {
	case config.DownloaderTypeAria2:
		form.Api = config.AppConfig.Downloader.Aria2.Api
		form.Token = config.AppConfig.Downloader.Aria2.Token
	case config.DownloaderTypeQBittorrent:
		form.Api = config.AppConfig.Downloader.QBittorrent.Api
		form.Username = config.AppConfig.Downloader.QBittorrent.Username
		form.Token = config.AppConfig.Downloader.QBittorrent.Password
	}
	api.OkResponse(c, form)
}
