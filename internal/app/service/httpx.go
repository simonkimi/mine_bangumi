package service

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/pkg/request"
	"github.com/simonkimi/minebangumi/tools/xstring"
)

type HttpX interface {
	GetClient() request.Client
}

type HttpxConfigImpl struct {
	config config.Config
}

func newHttpX(config config.Config) *HttpxConfigImpl {
	return &HttpxConfigImpl{config: config}
}

func (h *HttpxConfigImpl) GetClient() request.Client {
	enableProxy := h.config.GetBool(config.ProxyEnabled)
	if !enableProxy {
		return request.NewRestyClient(resty.New())
	}
	proxyScheme := h.config.GetString(config.ProxyScheme)
	proxyHost := h.config.GetString(config.ProxyHost)
	proxyPort := h.config.GetString(config.ProxyPort)
	proxyUsername := h.config.GetString(config.ProxyUsername)
	proxyPassword := h.config.GetString(config.ProxyPassword)
	c := resty.New().SetProxy(getProxyUrl(proxyScheme, proxyHost, proxyPort, proxyUsername, proxyPassword))
	return request.NewRestyClient(c)
}

func getProxyUrl(scheme string, host string, port string, username string, password string) string {
	if !xstring.IsEmptyOrWhitespace(username) {
		return fmt.Sprintf("%s://%s:%s@%s:%s", scheme, username, password, host, port)
	}
	return fmt.Sprintf("%s://%s:%s", scheme, host, port)
}
