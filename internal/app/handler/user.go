package handler

import (
	"encoding/base64"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"time"
)

func generateJwt(username string) (string, error) {
	expireTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &api.UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := base64.URLEncoding.DecodeString(config.AppConfig.System.SecretKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode secret key")
	}
	return token.SignedString(secret)
}

func Login(c *gin.Context) {
	var form api.LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		_ = c.AbortWithError(200, errno.NewFormError(err))
		return
	}
	if form.Username != config.AppConfig.User.Username || form.Password != config.AppConfig.User.Password {
		_ = c.AbortWithError(200, errno.NewApiError(errno.ErrorUserPasswordWrong))
		return
	}
	token, err := generateJwt(form.Username)
	if err != nil {
		_ = c.AbortWithError(200, err)
		return
	}
	c.JSON(200, &api.TokenResponse{
		Token: token,
	})
}
