package middleware

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service/user_service"
	"github.com/simonkimi/minebangumi/pkg/tests"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	cleanUp := tests.MainOnTempDir()
	defer cleanUp()
	config.Setup()
	m.Run()
}

func initAuthGin() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(JwtAuthMiddleware())
	return r
}

func TestJwtAuthMiddleware_Ok(t *testing.T) {
	username := "admin"
	token, err := user_service.GenerateJwt(username)
	require.Nil(t, err)

	r := initAuthGin()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"claims": GetClaims(c),
		})
	})

	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.Default(t, server.URL).
		GET("/test").
		WithHeader("Authorization", "Bearer "+token).
		Expect().
		Status(200).
		JSON().Object().
		Value("claims").Object().
		HasValue("username", username)
}

func TestJwtAuthMiddleware_NoToken(t *testing.T) {
	r := initAuthGin()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"claims": GetClaims(c),
		})
	})

	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.Default(t, server.URL).
		GET("/test").
		Expect().
		Status(200).
		JSON().Object().
		Value("claims").IsNull()
}

func TestJwtAuthMiddleware_TokenError(t *testing.T) {
	r := initAuthGin()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"claims": GetClaims(c),
		})
	})

	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.Default(t, server.URL).
		GET("/test").
		WithHeader("Authorization", "Bearer invalid-token").
		Expect().
		Status(200).
		JSON().Object().
		Value("claims").IsNull()
}
