package api

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
