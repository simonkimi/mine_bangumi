package middleware

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/pkg/tests"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	cleanUp := tests.MainOnTempDir()
	defer cleanUp()
	config.NewConfig()
	m.Run()
}

func TestJwtAuthMiddleware_Ok(t *testing.T) {
	username := "admin"
	token, err := service.GenerateUserJwt(username)
	require.Nil(t, err)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(JwtAuthMiddleware())
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
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(JwtAuthMiddleware())
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
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(JwtAuthMiddleware())
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

func TestRequireAuthMiddleware_Ok(t *testing.T) {
	//gin.SetMode(gin.TestMode)
	//r := gin.Default()
	//NewConfig(r)
	//group := r.Group("/api")
	//group.Use(RequireAuthMiddleware())
	//group.GET("/test", func(c *gin.Context) {
	//	api.OkResponse(c, &gin.H{
	//		"key": "value",
	//	})
	//})
	//
	//server := httptest.NewServer(r)
	//defer server.Close()
	//
	//username := "admin"
	//token, err := service.GenerateUserJwt(username)
	//require.Nil(t, err)
	//
	//httpexpect.Default(t, server.URL).
	//	GET("/api/test").
	//	WithHeader("Authorization", "Bearer "+token).
	//	Expect().
	//	Status(http.StatusOK).
	//	JSON().Object().
	//	HasValue("code", errno.Success).
	//	Value("data").Object().HasValue("key", "value")
}

func TestRequireAuthMiddleware_Error(t *testing.T) {
	//gin.SetMode(gin.TestMode)
	//r := gin.Default()
	//NewConfig(r)
	//group := r.Group("/api")
	//group.Use(RequireAuthMiddleware())
	//group.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"data": "ok",
	//	})
	//})
	//
	//server := httptest.NewServer(r)
	//defer server.Close()
	//
	//httpexpect.Default(t, server.URL).
	//	GET("/api/test").
	//	Expect().
	//	Status(http.StatusOK).
	//	JSON().Object().
	//	HasValue("code", errno.Unauthorized).
	//	HasValue("message", "Unauthorized")
}
