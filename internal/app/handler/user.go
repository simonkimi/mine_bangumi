package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
)

// Login godoc
// @Summary User login
// @Description Authenticate user and return token
// @Tags user
// @Accept json
// @Produce json
// @Param form body api.LoginForm true "Login Form"
// @Success 200 {object} api.TokenResponse "Token"
// @Router /api/v1/user/login [post]
func (w *WebApi) login(c *gin.Context) {
	var form api.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		_ = c.Error(api.NewBadRequestErrorf("invalid request: %s", err))
		c.Abort()
		return
	}

	username := w.mgr.GetConfig().GetString(config.UserUsername)
	password := w.mgr.GetConfig().GetString(config.UserPassword)
	if form.Username != username || form.Password != password {
		_ = c.Error(&api.Error{
			Message: "invalid username or password",
			Code:    api.APIStatusEnumUserCredentialsError,
		})
		return
	}
	token := w.mgr.GetConfig().GetString(config.UserApiToken)
	api.OkResponse(c, &api.TokenResponse{
		Token: token,
	})
}

func (w *WebApi) initUser(c *gin.Context) {
	var form api.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		_ = c.Error(api.NewBadRequestErrorf("invalid request: %s", err))
		c.Abort()
		return
	}
	conf := w.mgr.GetConfig()
	if conf.GetBool(config.SystemInit) {
		_ = c.Error(api.NewForbiddenError())
		c.Abort()
		return
	}
	token := config.UpdateUser(conf, &form.Username, &form.Password)
	api.OkResponse(c, &api.TokenResponse{
		Token: token,
	})
}
