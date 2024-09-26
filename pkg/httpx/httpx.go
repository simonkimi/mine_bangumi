package httpx

import (
	"github.com/go-resty/resty/v2"
	"github.com/simonkimi/minebangumi/tools/xnet"
	"sync"
)

type Config struct {
	ProxyEnabled  bool
	ProxyScheme   string
	ProxyHost     string
	ProxyPort     string
	ProxyUsername string
	ProxyPassword string
}

type HttpX struct {
	clients sync.Map
	config  *Config
}

func NewHttpX(config *Config) *HttpX {
	if config != nil {
		return &HttpX{
			config: config,
		}
	}
	return &HttpX{
		config: &Config{
			ProxyEnabled: false,
		},
	}
}

func (h *HttpX) setClientProxy(client *resty.Client) {
	if h.config.ProxyEnabled {
		client.SetProxy(xnet.GetProxyUrl(
			h.config.ProxyScheme,
			h.config.ProxyHost,
			h.config.ProxyPort,
			h.config.ProxyUsername,
			h.config.ProxyPassword,
		))
	}
}

func (h *HttpX) newHttpClient(baseUrl string) *resty.Client {
	client := resty.New()
	client.BaseURL = baseUrl
	if h.config.ProxyEnabled {
		h.setClientProxy(client)
	}
	return client
}

func (h *HttpX) GetClient(baseUrl string) *resty.Client {
	client, _ := h.clients.LoadOrStore(baseUrl, h.newHttpClient(baseUrl))
	return client.(*resty.Client)
}

func (h *HttpX) GetTempClient() *resty.Client {
	client := resty.New()
	h.setClientProxy(client)
	return client
}
