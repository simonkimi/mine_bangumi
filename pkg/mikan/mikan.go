package mikan

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/api"
	"strconv"
	"strings"
)

type Client struct {
	client *resty.Client
}

func NewClient(c *resty.Client) *Client {
	return &Client{
		client: c,
	}
}

func (c *Client) ParseBangumiByUrl(ctx context.Context, url string) (*Bangumi, error) {
	resp, err := c.client.R().SetContext(ctx).Get(url)
	if err != nil || resp.IsError() {
		if errors.As(err, context.Canceled) {
			return nil, api.NewCancelErrorf("Mikan feed request canceled: %s", url)
		}
		if errors.As(err, context.DeadlineExceeded) {
			return nil, api.NewTimeoutErrorf("Mikan feed request timeout: %s", url)
		}
		return nil, api.NewThirdPartyErrorf(err, url, "failed to fetch mikan feed")
	}
	feed := gofeed.NewParser()
	feedData, err := feed.ParseString(string(resp.Body()))
	if err != nil {
		return nil, api.NewThirdPartyErrorf(err, url, "failed to parse mikan feed: %s", url)
	}

	bangumi := &Bangumi{
		Title:    strings.TrimSpace(strings.Replace(feedData.Title, "Mikan Project -", "", -1)),
		Episodes: make([]*Episode, 0),
	}
	for _, item := range feedData.Items {
		torrentSize, err := strconv.Atoi(item.Enclosures[0].Length)
		if err != nil {
			torrentSize = 0
		}

		bangumi.Episodes = append(bangumi.Episodes, &Episode{
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
