package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/pkg/request"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
)

func GetPoster(ctx context.Context, client request.Client, targetType string, target string) ([]byte, error) {
	switch targetType {
	case api.ScraperEnumTmdb.String():
		url := fmt.Sprintf("%s/t/p/w780%s", tmdb.ImageHost, target)
		rsp := client.R().SetContext(ctx).Get(url)
		if err := rsp.Error(); err != nil {
			if errors.As(err, context.Canceled) {
				return nil, api.NewCancelError()
			}
			if errors.Is(err, context.DeadlineExceeded) {
				return nil, api.NewTimeoutErrorf("Failed to get poster: %s", target)
			}
			return nil, api.NewThirdPartyErrorf(err, url, "Failed to get poster: %s", target)
		}
		return rsp.Body(), nil
	default:
		return nil, api.NewBadRequestErrorf("Unsupported scraper: %s", targetType)
	}
}
