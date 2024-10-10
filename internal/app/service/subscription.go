package service

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/pkg/bangumi"
	"github.com/simonkimi/minebangumi/pkg/mikan"
)

func ParseAcgSubscriptionSource(ctx context.Context, client *resty.Client, targetUrl string, source api.SourceEnum) (*api.ParseAcgSubscriptionResult, error) {
	switch source {
	case api.SourceEnumBangumi:
		mikanResult, err := mikan.ParseBangumiUrl(ctx, client, targetUrl)
		if err != nil {
			return nil, err
		}
		return parseMikanResult(mikanResult), nil

	default:
		return nil, api.NewBadRequestErrorf("Unknown parser: %s", source.String())
	}
}

func parseMikanResult(result *mikan.Bangumi) *api.ParseAcgSubscriptionResult {
	bangumiTitle := result.Title
	season := -1
	if len(result.Episodes) != 0 {
		bf := bangumi.ParseBangumiSourceName(result.Episodes[0].Title, "")
		if bf != nil {
			bangumiTitle = bf.Title
			season = bf.Season
		}
	}

	var files []string
	for _, episode := range result.Episodes {
		files = append(files, episode.Title)
	}

	return &api.ParseAcgSubscriptionResult{
		Title:  bangumiTitle,
		Season: season,
		Files:  files,
	}
}

//func magnetParse(ctx context.Context, targetUrl string, rawData *parserRawData) error {
//	magnetData, exist := cache.Get[magnet.FileInfo](ParserMagnet, targetUrl)
//	if !exist {
//		newMagnetData, err := magnet.ParseMagnet(ctx, targetUrl)
//		if err != nil {
//			return err
//		}
//		cache.Add(ParserMagnet, targetUrl, newMagnetData)
//		magnetData = newMagnetData
//	}
//	rawData.Title = magnetData.Name
//	rawData.Files = magnetData.Files
//	return nil
//}
