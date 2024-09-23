package service

import (
	"encoding/base64"
	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"time"
)

func GenerateUserJwt(username string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := &api.UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := base64.URLEncoding.DecodeString(config.appConfig.System.SecretKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode secret key")
	}
	return token.SignedString(secret)
}
