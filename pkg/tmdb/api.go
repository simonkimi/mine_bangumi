package tmdb

import (
	"github.com/cockroachdb/errors"
	"github.com/simonkimi/minebangumi/domain"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/pkg/http_client"
	"github.com/simonkimi/minebangumi/tools/stringt"
	"strconv"
)

func getTmdbLanguage(language string) (string, error) {
	switch language {
	case domain.LanguageEn:
		return "en-US", nil
	case domain.LanguageZhHans:
		return "zh-CN", nil
	case domain.LanguageZhHant:
		return "zh-TW", nil
	case domain.LanguageJa:
		return "ja-JP", nil
	}
	return "", errors.Newf("Unsupported language: %s", language)
}

func getApiKey() string {
	if stringt.IsEmptyOrWhitespace(config.AppConfig.Tmdb.ApiKey) {
		return domain.TmdbDefaultApikey
	}
	return config.AppConfig.Tmdb.ApiKey
}

func Search(title string) ([]*SearchResultItem, error) {
	client := http_client.GetClient(domain.TmdbHost)
	page := 1
	results := make([]*SearchResultItem, 0)
	for {
		var result rawSearchResult
		req, err := client.R().
			SetQueryParams(map[string]string{
				"api_key":       getApiKey(),
				"page":          strconv.Itoa(page),
				"query":         title,
				"include_adult": "true",
			}).
			SetResult(&result).
			Get("/3/search/tv")
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to search tmdb tv series: %s", title)
		}
		if req.IsError() {
			return nil, errors.Newf("Failed to search tmdb tv series: %s, status code: %d", title, req.StatusCode())
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

func QueryForDetail(id int, language string) (*DetailData, error) {
	client := http_client.GetClient(domain.TmdbHost)
	language, err := getTmdbLanguage(language)
	if err != nil {
		return nil, errors.Wrapf(err, "GetDetail")
	}
	var detail DetailData
	req, err := client.R().
		SetQueryParams(map[string]string{
			"api_key":  getApiKey(),
			"language": language,
		}).
		SetResult(&detail).
		Get("/3/tv/" + strconv.Itoa(id))
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get tmdb tv series detail: %d", id)
	}
	if req.IsError() {
		return nil, errors.Newf("Failed to get tmdb tv series detail: %d, status code: %d", id, req.StatusCode())
	}
	return &detail, nil
}
