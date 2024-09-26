package tmdb

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/tools/xstring"
	"strconv"
)

type Config struct {
	apiKey string
	client func() *resty.Client
}

func NewConfig(apiKey string, client func() *resty.Client) *Config {
	return &Config{apiKey: apiKey, client: client}
}

type Tmdb struct {
	config *Config
}

func NewTmdb(config *Config) *Tmdb {
	return &Tmdb{config: config}
}

func (t *Tmdb) GetTmdbLanguage(language api.ScraperLanguage) (string, error) {
	switch language {
	case api.ScraperLanguageEn:
		return "en-US", nil
	case api.ScraperLanguageZhHans:
		return "zh-CN", nil
	case api.ScraperLanguageZhHant:
		return "zh-TW", nil
	case api.ScraperLanguageJa:
		return "ja-JP", nil
	}
	return "", errors.New(fmt.Sprintf("Unsupported language: %s", language))
}

func (t *Tmdb) getApiKey() string {
	if xstring.IsEmptyOrWhitespace(t.config.apiKey) {
		return api.TmdbDefaultApikey
	}
	return t.config.apiKey
}

func (t *Tmdb) Search(ctx context.Context, title string) ([]*SearchResultItem, error) {
	page := 1
	results := make([]*SearchResultItem, 0)
	for {
		var result rawSearchResult
		req, err := t.config.client().R().
			SetContext(ctx).
			SetQueryParams(map[string]string{
				"api_key":       t.getApiKey(),
				"page":          strconv.Itoa(page),
				"query":         title,
				"include_adult": "true",
			}).
			SetResult(&result).
			Get("/3/search/tv")
		if err != nil {
			if errors.As(err, context.Canceled) {
				return nil, api.NewCancelErrorf("Search tmdb tv series canceled: %s", title)
			}
			if errors.As(err, context.DeadlineExceeded) {
				return nil, api.NewTimeoutErrorf("Search tmdb tv series timeout: %s", title)
			}
			return nil, api.NewThirdPartyErrorf(err, req.Request.URL, "Failed to search tmdb tv series: %s", title)
		}

		results = append(results, result.Results...)
		if result.Page < result.TotalPages {
			page = result.Page + 1
			continue
		}
		break
	}
	return results, nil
}

func (t *Tmdb) QueryForDetail(ctx context.Context, id int, language string) (*DetailData, error) {
	var detail DetailData
	req, err := t.config.client().R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"api_key":  t.getApiKey(),
			"language": language,
		}).
		SetResult(&detail).
		Get("/3/tv/" + strconv.Itoa(id))
	if err != nil || req.IsError() {
		return nil, api.NewThirdPartyErrorf(err, req.Request.URL, "Failed to get tmdb tv series detail: %d", id)
	}
	return &detail, nil
}
