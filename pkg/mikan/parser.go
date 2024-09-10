package mikan

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/simonkimi/minebangumi/pkg/http_client"
	"strconv"
)

func ParseFromUrl(url string) (*Bangumi, error) {
	client := http_client.GetTempClient()
	resp, err := client.R().Get(url)
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

	bangumi := &Bangumi{
		Title:    feedData.Title,
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
