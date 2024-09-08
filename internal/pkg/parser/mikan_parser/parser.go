package mikan_parser

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/mmcdole/gofeed"
	"github.com/simonkimi/minebangumi/internal/pkg/config"
	"github.com/simonkimi/minebangumi/tools/nett"
	"strconv"
	"sync"
)

var mikanParser *MikanParser
var once sync.Once

type MikanParser struct {
	client *resty.Client
}

func getMikanParser() *MikanParser {
	once.Do(func() {
		client := resty.New()
		if config.AppConfig.MikanProxy.Enable {
			client.SetProxy(nett.GetProxyUrl(
				config.AppConfig.MikanProxy.Scheme,
				config.AppConfig.MikanProxy.Host,
				config.AppConfig.MikanProxy.Port,
				config.AppConfig.MikanProxy.UseAuth,
				config.AppConfig.MikanProxy.Username,
				config.AppConfig.MikanProxy.Password,
			))
		}
		mikanParser = &MikanParser{
			client: client,
		}
	})
	return mikanParser
}

func ParseFromUrl(url string) (*MikanBangumi, error) {
	p := getMikanParser()
	resp, err := p.client.R().Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get url: %w", err)
	}
	if resp.StatusCode() != 200 {
		return nil, nil
	}

	feed := gofeed.NewParser()
	feedData, err := feed.ParseString(string(resp.Body()))
	if err != nil {
		return nil, fmt.Errorf("failed to parse mikan feed: %w", err)
	}

	bangumi := &MikanBangumi{
		Title:    feedData.Title,
		Episodes: make([]*MikanEpisode, 0),
	}
	for _, item := range feedData.Items {
		torrentSize, err := strconv.Atoi(item.Enclosures[0].Length)
		if err != nil {
			torrentSize = 0
		}

		bangumi.Episodes = append(bangumi.Episodes, &MikanEpisode{
			Title:       item.Title,
			Guid:        item.GUID,
			Link:        item.Link,
			Description: item.Description,
			Torrent:     item.Enclosures[0].URL,
			TorrentSize: torrentSize,
		})
	}
	return bangumi, nil
}

func ReloadConfig() {
	p := getMikanParser()
	if config.AppConfig.MikanProxy.Enable {
		p.client.SetProxy(nett.GetProxyUrl(
			config.AppConfig.MikanProxy.Scheme,
			config.AppConfig.MikanProxy.Host,
			config.AppConfig.MikanProxy.Port,
			config.AppConfig.MikanProxy.UseAuth,
			config.AppConfig.MikanProxy.Username,
			config.AppConfig.MikanProxy.Password,
		))
	} else {
		p.client.RemoveProxy()
	}
}
