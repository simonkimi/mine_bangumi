package tmdb

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/simonkimi/minebangumi/api"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch(t *testing.T) {
	result, err := Search(context.Background(), resty.New(), DefaultApikey, "魔法禁书目录")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestDetail(t *testing.T) {
	language, err := GetTmdbLanguage(api.ScraperLanguageEn)
	require.Nil(t, err)
	result, err := QueryForDetail(context.Background(), resty.New(), DefaultApikey, 30980, language)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
