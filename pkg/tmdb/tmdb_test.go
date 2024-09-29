package tmdb

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/simonkimi/minebangumi/api"
	"github.com/stretchr/testify/require"
	"testing"
)

func newTestTmdb() *Tmdb {
	return NewTmdb(NewConfig(api.TmdbDefaultApikey, func() *resty.Client {
		return resty.New()
	}))
}

func TestSearch(t *testing.T) {
	result, err := newTestTmdb().Search(context.Background(), "魔法禁书目录")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestDetail(t *testing.T) {
	language, err := GetTmdbLanguage(api.ScraperLanguageEn)
	require.Nil(t, err)
	result, err := newTestTmdb().QueryForDetail(context.Background(), 30980, language)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
