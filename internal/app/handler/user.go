package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/pkg/errno"
)

// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param form body api.LoginForm true "Login Form"
// @Success 200 {object} api.TokenResponse "JWT Token"
// @Router /api/v1/user/login [post]
func Login(c *gin.Context) {
	var form api.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		_ = c.Error(errno.NewFormError(err))
		return
	}
	if form.Username != config.appConfig.User.Username || form.Password != config.appConfig.User.Password {
		_ = c.Error(errno.NewApiError(errno.ErrorUserPasswordWrong))
		return
	}
	token, err := service.GenerateUserJwt(form.Username)
	if err != nil {
		_ = c.Error(err)
		return
	}
	api.OkResponse(c, &api.TokenResponse{
		Token: token,
	})
}
