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

const (
	DefaultApikey = "32b19d6a05b512190a056fa4e747cbbc"
	Host          = "https://api.themoviedb.org"
	ImageHost     = "https://image.tmdb.org"
)

func GetTmdbLanguage(language api.ScraperLanguage) (string, error) {
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

func getApiKey(key string) string {
	if xstring.IsEmptyOrWhitespace(key) {
		return DefaultApikey
	}
	return key
}

func Search(ctx context.Context, client *resty.Client, apiKey string, title string) ([]*SearchResultItem, error) {
	page := 1
	results := make([]*SearchResultItem, 0)
	for {
		var result rawSearchResult
		req, err := client.
			SetBaseURL(Host).
			R().
			SetContext(ctx).
			SetQueryParams(map[string]string{
				"api_key":       getApiKey(apiKey),
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

func QueryForDetail(ctx context.Context, client *resty.Client, apiKey string, id int, language string) (*DetailData, error) {
	var detail DetailData
	req, err := client.
		SetBaseURL(Host).
		R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"api_key":  getApiKey(apiKey),
			"language": language,
		}).
		SetResult(&detail).
		Get("/3/tv/" + strconv.Itoa(id))
	if err != nil || req.IsError() {
		return nil, api.NewThirdPartyErrorf(err, req.Request.URL, "Failed to get tmdb tv series detail: %d", id)
	}
	return &detail, nil
}

type rawSearchResult struct {
	Page         int                 `json:"page"`
	Results      []*SearchResultItem `json:"results"`
	TotalPages   int                 `json:"total_pages"`
	TotalResults int                 `json:"total_results"`
}
