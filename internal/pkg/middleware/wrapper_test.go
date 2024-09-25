package middleware

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initGin() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(ResponseWrapperMiddleware())
	return r
}

func TestResponseWrapperMiddleware_Ok(t *testing.T) {
	r := initGin()
	r.GET("/test", func(c *gin.Context) {
		api.OkResponse(c, gin.H{
			"key": "value",
		})
	})
	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.Default(t, server.URL).
		GET("/test").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("code").HasValue("code", errno.Success).
		ContainsKey("data").HasValue("data", gin.H{
		"key": "value",
	})
}

func TestResponseWrapperMiddleware_ApiError(t *testing.T) {
	r := initGin()
	r.GET("/test", func(c *gin.Context) {
		_ = c.Error(errno.NewApiErrorf(errno.ErrorApiNetwork, "network error"))
	})
	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.Default(t, server.URL).
		GET("/test").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("code").HasValue("code", errno.ErrorApiNetwork).
		Value("message").String().NotEmpty()
}

func TestResponseWrapperMiddleware_Error(t *testing.T) {
	r := initGin()
	r.GET("/test", func(c *gin.Context) {
		_ = c.Error(errors.New("internal error"))
	})
	server := httptest.NewServer(r)
	defer server.Close()

	httpexpect.Default(t, server.URL).
		GET("/test").
		Expect().
		Status(http.StatusInternalServerError).
		JSON().Object().
		ContainsKey("code").HasValue("code", errno.InternalServerError).
		Value("message").String().NotEmpty()
}
