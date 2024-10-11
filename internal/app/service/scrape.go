package service

import (
	"context"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/pkg/request"
	"github.com/simonkimi/minebangumi/pkg/tmdb"
	"strconv"
)

func ScrapeTmDb(ctx context.Context, client request.Client, apiKey string, form *api.ScrapeSearchInput) ([]*api.ScrapeSearchResult, error) {
	language, err := tmdb.GetTmdbLanguage(form.Language)
	if err != nil {
		return nil, err
	}

	var result []*api.ScrapeSearchResult
	tmdbSearch, err := tmdb.Search(ctx, client, apiKey, form.Title)
	if err != nil {
		return nil, err
	}
	for _, item := range tmdbSearch {
		detail, err := tmdb.QueryForDetail(ctx, client, apiKey, item.ID, language)
		if err != nil {
			return nil, err
		}

		var seasons []*api.ScrapeSearchSeasonResult
		for _, season := range detail.Seasons {
			seasons = append(seasons, &api.ScrapeSearchSeasonResult{
				SeasonID: season.ID,
				Title:    season.Name,
				Overview: season.Overview,
				Poster:   season.PosterPath,
			})
		}

		result = append(result, &api.ScrapeSearchResult{
			Scraper:       api.ScraperEnumTmdb,
			ID:            strconv.Itoa(detail.ID),
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
