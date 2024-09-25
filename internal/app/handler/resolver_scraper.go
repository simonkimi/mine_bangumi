package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/api"
)

// ScraperSource is the resolver for the scraperSource field.
func (r *queryResolver) ScraperSource(ctx context.Context, input api.ParseAcgSourceInput) (*api.ParseAcgSourceResult, error) {
	//return service.ParseSource(ctx, input.Source, input.Parser)
	return r.mgr.Source.ParseSource(ctx, input.Source, input.Parser)
}

// ScraperDb is the resolver for the scraperDb field.
func (r *queryResolver) ScraperDb(ctx context.Context, input api.ScrapeAcgSourceInput) ([]*api.ScrapeAcgResult, error) {
	return r.mgr.Scraper.ScrapeService(ctx, &input)
}
