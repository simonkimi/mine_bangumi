package service

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/api"
)

func GetPoster(ctx context.Context, client *resty.Client, targetType string, target string) ([]byte, error) {
	switch targetType {
	case ScrapeTmDb:
		url := fmt.Sprintf("%s/t/p/w780%s", api.TmdbImageHost, target)
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
			return nil, api.NewBadRequestErrorf("Failed to get poster: %s, status code: %d", target, rsp.StatusCode())
		}
		return rsp.Body(), nil
	default:
		return nil, api.NewBadRequestErrorf("Unsupported scraper: %s", targetType)
	}
}
