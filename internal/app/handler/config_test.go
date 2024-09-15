package handler

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/domain"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/pkg/tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	gin.SetMode(gin.TestMode)
	tests.MainOnTempDir()
	config.Setup()
}

func TestSystem(t *testing.T) {
	r := gin.Default()
	r.GET("/system", System)
	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.
		Default(t, server.URL).
		GET("/system").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		HasValue("code", 200).
		Value("data").Object().
		HasValue("version", domain.Version).
		HasValue("is_first_run", true).
		HasValue("is_login", false)
}

func TestInitUser(t *testing.T) {
	r := gin.Default()
	r.POST("/init_user", InitUser)
	server := httptest.NewServer(r)
	defer server.Close()

	username := "root"
	password := "123456"

	httpexpect.
		Default(t, server.URL).
		POST("/init_user").
		WithJSON(gin.H{
			"username": username,
			"password": password,
		}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		HasValue("code", 200).
		Value("data").Object().
		Value("token").String().NotEmpty()

	assert.Equal(t, config.AppConfig.User.Username, username)
	assert.Equal(t, config.AppConfig.User.Password, password)
}
