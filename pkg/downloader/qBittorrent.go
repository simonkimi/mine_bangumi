package downloader

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type QBittorrentClient struct {
	IDownloader
	url      string
	username string
	password string
	client   *resty.Client
}

func NewQBittorrentDl(url string, username string, password string) *QBittorrentClient {
	return &QBittorrentClient{
		url:      url,
		username: username,
		password: password,
		client:   resty.New().SetBaseURL(url),
	}
}

func (d *QBittorrentClient) Login() error {
	rsp, err := d.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"username": d.username,
			"password": d.password,
		}).
		Post("/api/v2/auth/login")

	if err != nil {
		return err
	}

	if rsp.IsError() {
		return fmt.Errorf("qBittorrentDl login failed: %s", rsp.Status())
	}
	cookie := rsp.Header().Get("Set-Cookie")
	if cookie == "" {
		return fmt.Errorf("qBittorrentDl login failed: no cookie")
	}
	logrus.Debugf("qBittorrentDl login success: %s", cookie)
	return nil
}

func (d *QBittorrentClient) RecordClientInfo() (string, error) {
	rsp, err := d.client.R().Get("/api/v2/app/version")
	if err != nil {
		return "", fmt.Errorf("get qb version failed: %s", err)
	}
	appVersion := rsp.String()
	rsp, err = d.client.R().Get("/api/v2/app/webapiVersion")
	if err != nil {
		return "", fmt.Errorf("get qb webapi version failed: %s", err)
	}
	webApiVersion := rsp.String()
	version := fmt.Sprintf("qBittorrent version: %s, webapi version: %s", appVersion, webApiVersion)
	logrus.Debugf("qBittorrent version: %s", version)
	return version, nil

}
