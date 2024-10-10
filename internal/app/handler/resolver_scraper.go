package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/internal/app/service"
)

func (q *queryResolver) ScraperSearch(ctx context.Context, input api.ScrapeSearchInput) ([]*api.ScrapeSearchResult, error) {
	switch input.Scraper {
	case api.ScraperEnumTmdb:
		apiKey := q.mgr.GetConfig().GetString(config.TmdbApiKey)
		client := q.mgr.GetHttpX().GetClient()
		return service.ScrapeTmDb(ctx, client, apiKey, &input)
	default:
		return nil, api.NewBadRequestErrorf("unsupported scraper %s", input.Scraper)
	}
}
