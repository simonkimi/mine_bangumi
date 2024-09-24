package service

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
)

const ScrapeTmDb = "tmdb"

func ScrapeService(ctx context.Context, input *api.ScrapeAcgSourceInput) ([]*api.ScrapeAcgResult, error) {
	switch input.Scraper {
	case api.ScraperEnumTmdb:
		return scrapeTmDb(ctx, input)
	default:
		return nil, errors.Newf("Unsupported scraper: %s", input.Scraper)
	}
}

func scrapeTmDb(ctx context.Context, form *api.ScrapeAcgSourceInput) ([]*api.ScrapeAcgResult, error) {
	language, err := tmdb.GetTmdbLanguage(form.Language)
	if err != nil {
		return nil, err
	}

	var result []*api.ScrapeAcgResult
	tmdbSearch, err := tmdb.Search(ctx, form.Title)
	if err != nil {
		return nil, err
	}
	for _, item := range tmdbSearch {
		detail, err := tmdb.QueryForDetail(ctx, item.ID, language)
		if err != nil {
			return nil, err
		}

		var seasons []*api.ScrapeAcgSeasonResult
		for _, season := range detail.Seasons {
			seasons = append(seasons, &api.ScrapeAcgSeasonResult{
				SeasonID: season.ID,
				Title:    season.Name,
				Overview: season.Overview,
				Poster:   season.PosterPath,
			})
		}

		result = append(result, &api.ScrapeAcgResult{
			Scraper:       ScrapeTmDb,
			Title:         detail.Name,
			OriginalTitle: detail.OriginalName,
			FirstAirDate:  detail.FirstAirDate,
			Overview:      detail.Overview,
			Backdrop:      detail.BackdropPath,
			Poster:        detail.PosterPath,
			Seasons:       seasons,
		})
	}
	return result, nil
}
