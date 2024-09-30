package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
)

func (w *WebApi) systemStatus(c *gin.Context) {
	db := w.mgr.GetDatabase()
	conf := w.mgr.GetConfig()

	isLogin := IsLogin(c)
	username := ""
	if isLogin {
		username = conf.GetString(config.UserUsername)
	}

	api.OkResponse(c, &api.SystemInfo{
		Version:                api.Version,
		AppDatabaseVersion:     db.GetAppSchemaVersion(),
		CurrentDatabaseVersion: db.GetSchemaVersion(),
		IsSystemInit:           conf.GetBool(config.SystemInit),
		IsLogin:                isLogin,
		Username:               username,
	})
}
