package handler

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/pkg/middleware"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"github.com/simonkimi/minebangumi/pkg/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initGin() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	middleware.Setup(r)
	return r
}

func TestMain(m *testing.M) {
	clean := tests.MainOnTempDir()
	config.Setup()
	defer clean()
	m.Run()
}

func TestLogin_Success(t *testing.T) {
	r := initGin()
	r.POST("/login", Login)
	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.POST("/login").
		WithJSON(&api.LoginForm{
			Username: config.AppConfig.User.Username,
			Password: config.AppConfig.User.Password,
		}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("code").HasValue("code", errno.Success).
		Value("data").Object().
		Value("token").String().NotEmpty()
}

func TestLogin_InvalidCredentials(t *testing.T) {
	r := initGin()
	r.POST("/login", Login)

	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.POST("/login").
		WithJSON(&api.LoginForm{
			Username: config.AppConfig.User.Username + "invalid",
			Password: config.AppConfig.User.Password + "invalid",
		}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("code").HasValue("code", errno.ErrorUserPasswordWrong).
		Value("message").String().NotEmpty()
}

func TestLogin_InvalidJson(t *testing.T) {
	r := initGin()
	r.POST("/login", Login)

	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.POST("/login").
		WithText("data").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("code").HasValue("code", errno.BadRequest).
		Value("message").String().NotEmpty()
}
