package service

import (
	"context"
	"github.com/simonkimi/minebangumi/domain"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScrapeService(t *testing.T) {
	form := &api.ScrapeForm{
		Scraper:  ScrapeTmDb,
		Title:    "不死者之王",
		Language: domain.LanguageZhHans,
	}
	result, err := ScrapeService(context.Background(), form)
	require.Nil(t, err)
	t.Log(result)
}
