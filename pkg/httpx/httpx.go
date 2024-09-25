package httpx

import (
	"github.com/go-resty/resty/v2"
	"github.com/simonkimi/minebangumi/tools/xnet"
	"sync"
)

type HttpX struct {
	clients       sync.Map
	proxyEnabled  bool
	proxyScheme   string
	proxyHost     string
	proxyPort     string
	proxyUsername string
	proxyPassword string
}

func NewHttpX(proxyEnabled bool, proxyScheme, proxyHost, proxyPort, proxyUsername, proxyPassword string) *HttpX {
	return &HttpX{
		proxyEnabled:  proxyEnabled,
		proxyScheme:   proxyScheme,
		proxyHost:     proxyHost,
		proxyPort:     proxyPort,
		proxyUsername: proxyUsername,
		proxyPassword: proxyPassword,
	}
}

func (h *HttpX) setClientProxy(client *resty.Client) {
	if h.proxyEnabled {
		client.SetProxy(xnet.GetProxyUrl(
			h.proxyScheme,
			h.proxyHost,
			h.proxyPort,
			true,
			h.proxyUsername,
			h.proxyPassword,
		))
	}
}

func (h *HttpX) newHttpClient(baseUrl string) *resty.Client {
	client := resty.New()
	client.BaseURL = baseUrl
	if h.proxyEnabled {
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
