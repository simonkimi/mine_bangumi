package http_client

import (
	"github.com/go-resty/resty/v2"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/tools/nett"
	"sync"
)

var (
	clientContainer *httpClientManager
	clientOnce      sync.Once
)

type httpClientManager struct {
	clients sync.Map
}

func setClientProxy(client *resty.Client) {
	if config.appConfig.ProxyConfig.Enable {
		client.SetProxy(nett.GetProxyUrl(
			config.appConfig.ProxyConfig.Scheme,
			config.appConfig.ProxyConfig.Host,
			config.appConfig.ProxyConfig.Port,
			config.appConfig.ProxyConfig.UseAuth,
			config.appConfig.ProxyConfig.Username,
			config.appConfig.ProxyConfig.Password,
		))
	}
}

func newHttpClient(baseUrl string) *resty.Client {
	client := resty.New()
	client.BaseURL = baseUrl
	if config.appConfig.ProxyConfig.Enable {
		setClientProxy(client)
	}
	return client
}

func ReloadConfig() {
	if clientContainer == nil {
		return
	}
	clientContainer.clients.Range(func(key, value any) bool {
		client := value.(*resty.Client)
		if config.appConfig.ProxyConfig.Enable {
			setClientProxy(client)
		} else {
			client.RemoveProxy()
		}
		return true
	})
}

func GetClient(baseUrl string) *resty.Client {
	clientOnce.Do(func() {
		clientContainer = &httpClientManager{}
	})
	client, _ := clientContainer.clients.LoadOrStore(baseUrl, newHttpClient(baseUrl))
	return client.(*resty.Client)
}

func GetTempClient() *resty.Client {
	client := resty.New()
	setClientProxy(client)
	return client
}
