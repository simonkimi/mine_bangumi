package service

import (
	"context"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"github.com/simonkimi/minebangumi/pkg/http_client"
)

func GetPoster(ctx context.Context, targetType string, target string) ([]byte, error) {
	client := http_client.GetTempClient()
	switch targetType {
	case ScrapeTmDb:
		url := "https://image.tmdb.org/t/p/w780" + target
		rsp, err := client.R().SetContext(ctx).Get(url)
		if err != nil {
			return nil, errno.NewApiErrorWithCausef(errno.ErrorApiNetwork, err, "Failed to get poster: %s", target)
		}
		if rsp.IsError() {
			return nil, errno.NewApiErrorf(errno.ErrorApiParse, "Failed to get poster: %s, status code: %d", target, rsp.StatusCode())
		}
		return rsp.Body(), nil
	default:
		return nil, errno.NewApiErrorf(errno.BadRequest, "Unsupported scraper: %s", targetType)
	}
}
