package downloader

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type QBittorrentDl struct {
	IDownloader
	url      string
	username string
	password string
	client   *resty.Client
}

func NewQBittorrentDl(url string, username string, password string) *QBittorrentDl {
	return &QBittorrentDl{
		url:      url,
		username: username,
		password: password,
		client: resty.New().
			SetBaseURL(url),
	}
}

func (d *QBittorrentDl) Login() error {
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

	if rsp.StatusCode() != 200 {
		return fmt.Errorf("qb login failed: %s", rsp.Status())
	}
	logrus.Debugf("qb login success: %s", rsp.Header().Get("Set-Cookie"))
	return nil
}

func (d *QBittorrentDl) RecordClientInfo() (string, error) {
	rsp, err := d.client.R().Get("/api/v2/app/version")
	if err != nil {
		return "", fmt.Errorf("get qb version failed: %s", err)
	}
	appVersion := rsp.String()
	logrus.Infof("qb version: %s", rsp.String())
	rsp, err = d.client.R().Get("/api/v2/app/webapiVersion")
	if err != nil {
		return "", fmt.Errorf("get qb webapi version failed: %s", err)
	}
	webApiVersion := rsp.String()
	logrus.Infof("qb webapi version: %s", rsp.String())
	return fmt.Sprintf("qb version: %s, webapi version: %s", appVersion, webApiVersion), nil
}
