package handler

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/pkg/hash"
	"github.com/simonkimi/minebangumi/pkg/testutil"
	"github.com/simonkimi/minebangumi/tools/xstring"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testApiKey = hash.GenerateRandomKey(40)
var testUsername = xstring.RandomString(5)
var testPassword = xstring.RandomString(40)

func initApi() (*gin.Engine, *WebApi) {
	gin.SetMode(gin.TestMode)
	mgr := new(service.MockManager)
	webapi := NewWebApi(NewWebApiConfig(nil, mgr))
	conf := config.NewMockConfig()
	mgr.EXPECT().GetConfig().Return(conf)

	conf.SetString(config.UserUsername, testUsername)
	conf.SetString(config.UserPassword, testPassword)
	conf.SetString(config.UserApiToken, testApiKey)
	return webapi.Engine, webapi
}

func TestLogin_Success(t *testing.T) {
	r, web := initApi()
	r.POST("/login", web.login)
	server := httptest.NewServer(r)
	defer server.Close()

	e := testutil.NewDebugHttp(t, server.URL)
	e.POST("/login").
		WithJSON(&api.LoginForm{
			Username: testUsername,
			Password: testPassword,
		}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("code").HasValue("code", api.APIStatusEnumSuccess).
		Value("data").Object().
		Value("token").String().NotEmpty()
}

func TestLogin_InvalidCredentials(t *testing.T) {
	r, web := initApi()
	r.POST("/login", web.login)

	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.POST("/login").
		WithJSON(&api.LoginForm{
			Username: testUsername + "invalid",
			Password: testPassword + "invalid",
		}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("code").HasValue("code", api.APIStatusEnumUnauthorized).
		Value("message").String().NotEmpty()
}

func TestLogin_InvalidJson(t *testing.T) {
	r, web := initApi()
	r.POST("/login", web.login)

	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.POST("/login").
		WithText("data").
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("code").HasValue("code", api.APIStatusEnumBadRequest).
		Value("message").String().NotEmpty()
}

func TestInitUser_Ok(t *testing.T) {
	r, web := initApi()
	conf := web.mgr.GetConfig()
	conf.SetBool(config.SystemInit, false)

	r.POST("/init", web.initUser)
	server := httptest.NewServer(r)
	defer server.Close()

	testutil.NewDebugHttp(t, server.URL).
		POST("/init").
		WithJSON(&api.LoginForm{
			Username: testUsername,
			Password: testPassword}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("code").HasValue("code", api.APIStatusEnumSuccess).
		Value("data").Object().
		Value("token").String().NotEmpty()
}

func TestInitUser_Forbidden(t *testing.T) {
	r, web := initApi()
	conf := web.mgr.GetConfig()
	conf.SetBool(config.SystemInit, true)

	r.POST("/init", web.initUser)
	server := httptest.NewServer(r)
	defer server.Close()

	testutil.NewDebugHttp(t, server.URL).
		POST("/init").
		WithJSON(&api.LoginForm{
			Username: testUsername,
			Password: testPassword}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("code").HasValue("code", api.APIStatusEnumForbidden)
}
