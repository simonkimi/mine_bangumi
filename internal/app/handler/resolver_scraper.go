package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/service"
)

// ScraperSource is the resolver for the scraperSource field.
func (r *queryResolver) ScraperSource(ctx context.Context, input api.ParseAcgSourceInput) (*api.ParseAcgSourceResult, error) {
	return service.ParseSource(ctx, input.Source, input.Parser)
}

// ScraperDb is the resolver for the scraperDb field.
func (r *queryResolver) ScraperDb(ctx context.Context, input api.ScrapeAcgSourceInput) ([]*api.ScrapeAcgResult, error) {
	return service.ScrapeService(ctx, &input)
}
