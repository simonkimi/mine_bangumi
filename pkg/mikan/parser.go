package mikan

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/mmcdole/gofeed"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"github.com/simonkimi/minebangumi/pkg/http_client"
	"strconv"
	"strings"
)

func ParseBangumiByUrl(ctx context.Context, url string) (*Bangumi, error) {
	client := http_client.GetTempClient()
	resp, err := client.R().SetContext(ctx).Get(url)
	if err != nil {
		if errors.As(err, context.Canceled) {
			return nil, errno.NewApiError(errno.ErrorCancel)
		}
		if errors.As(err, context.DeadlineExceeded) {
			return nil, errno.NewApiError(errno.ErrorTimeout)
		}
		return nil, errno.NewApiErrorWithCause(errno.ErrorApiNetwork, err)
	}
	if resp.StatusCode() != 200 {
		return nil, errno.NewApiErrorf(errno.ErrorApiNetwork, "failed to fetch mikan feed, status code: %d", resp.StatusCode())
	}
	feed := gofeed.NewParser()
	feedData, err := feed.ParseString(string(resp.Body()))
	if err != nil {
		return nil, errno.NewApiErrorWithCausef(errno.ErrorApiParse, err, "failed to parse mikan feed: %s", url)
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
