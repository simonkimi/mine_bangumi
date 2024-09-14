package service

import (
	"context"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
)

const ScrapeTmDb = "tmdb"

func ScrapeService(ctx context.Context, form *api.ScrapeForm) ([]*api.ScrapeResponse, error) {
	switch form.Scraper {
	case ScrapeTmDb:
		return scrapeTmDb(ctx, form)
	default:
		return nil, errno.NewApiErrorf(errno.BadRequest, "Unsupported scraper: %s", form.Scraper)
	}
}

func scrapeTmDb(ctx context.Context, form *api.ScrapeForm) ([]*api.ScrapeResponse, error) {
	language, err := tmdb.GetTmdbLanguage(form.Language)
	if err != nil {
		return nil, err
	}

	var result []*api.ScrapeResponse
	tmdbSearch, err := tmdb.Search(ctx, form.Title)
	if err != nil {
		return nil, err
	}
	for _, item := range tmdbSearch {
		detail, err := tmdb.QueryForDetail(ctx, item.ID, language)
		if err != nil {
			return nil, err
		}

		var seasons []*api.ScrapeSeasonResponse
		for _, season := range detail.Seasons {
			seasons = append(seasons, &api.ScrapeSeasonResponse{
				SeasonId: season.ID,
				Title:    season.Name,
				Overview: season.Overview,
				Poster:   season.PosterPath,
			})
		}

		result = append(result, &api.ScrapeResponse{
			Scraper:       ScrapeTmDb,
			Title:         detail.Name,
			OriginalTitle: detail.OriginalName,
			FirstAirYear:  detail.FirstAirDate,
			Overview:      detail.Overview,
			PosterPath:    detail.PosterPath,
			Seasons:       seasons,
		})
	}
	return result, nil
}
