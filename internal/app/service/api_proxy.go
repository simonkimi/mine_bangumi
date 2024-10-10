package service

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
)

func GetPoster(ctx context.Context, client *resty.Client, targetType string, target string) ([]byte, error) {
	switch targetType {
	case api.ScraperEnumTmdb.String():
		url := fmt.Sprintf("%s/t/p/w780%s", tmdb.ImageHost, target)
		rsp, err := client.R().SetContext(ctx).Get(url)
		if err != nil {
			if errors.As(err, context.Canceled) {
				return nil, api.NewCancelError()
			}
			if errors.Is(err, context.DeadlineExceeded) {
				return nil, api.NewTimeoutErrorf("Failed to get poster: %s", target)
			}
			return nil, api.NewThirdPartyErrorf(err, url, "Failed to get poster: %s", target)
		}
		if rsp.IsError() {
			return nil, api.NewThirdPartyErrorf(nil, target, "Failed to get poster, status code: %d", rsp.StatusCode())
		}
		return rsp.Body(), nil
	default:
		return nil, api.NewBadRequestErrorf("Unsupported scraper: %s", targetType)
	}
}
