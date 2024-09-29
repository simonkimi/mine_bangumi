package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
)

type ApiProxyService struct {
	httpx *HttpX
}

func newApiProxyService(httpx *HttpX) *ApiProxyService {
	return &ApiProxyService{
		httpx: httpx,
	}
}

func (s *ApiProxyService) GetPoster(ctx context.Context, targetType string, target string) ([]byte, error) {
	switch targetType {
	case ScrapeTmDb:
		url := fmt.Sprintf("%s/t/p/w780%s", tmdb.TmdbImageHost, target)
		rsp, err := s.httpx.GetTempClient().R().SetContext(ctx).Get(url)
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
