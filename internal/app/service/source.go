package service

import (
	"context"
	"github.com/simonkimi/minebangumi/api"
	"github.com/simonkimi/minebangumi/internal/pkg/cache"
	"github.com/simonkimi/minebangumi/pkg/bangumi"
	"github.com/simonkimi/minebangumi/pkg/mikan"
)

const (
	ParserMikan = "mikan"
)

type SourceService struct {
	mikan *mikan.Client
	cache *cache.Cache
}

func newSourceService(mikan *mikan.Client) *SourceService {
	return &SourceService{mikan: mikan}
}

type RssModel struct {
	RssTitle string
}

type parserRawData struct {
	Title string
	Files []string
}

func (c *SourceService) ParseSource(ctx context.Context, targetUrl string, parser api.SourceParserEnum) (*api.ParseAcgSourceResult, error) {
	rawData := &parserRawData{}
	switch parser {
	case api.SourceParserEnumBangumi:
		err := c.mikanParse(ctx, targetUrl, rawData)
		if err != nil {
			return nil, err
		}
	default:
		return nil, api.NewBadRequestErrorf("Unknown parser: %s", parser.String())
	}
	bangumiTitle := rawData.Title
	season := -1
	if len(rawData.Files) != 0 {
		bf := bangumi.ParseBangumiSourceName(rawData.Files[0], "")
		if bf != nil {
			bangumiTitle = bf.Title
			season = bf.Season
		}
	}

	return &api.ParseAcgSourceResult{
		Title:  bangumiTitle,
		Season: season,
		Files:  rawData.Files,
	}, nil
}

func (c *SourceService) mikanParse(ctx context.Context, targetUrl string, rawData *parserRawData) error {
	var mikanData *mikan.Bangumi
	data, exist := c.cache.Get(ParserMikan, targetUrl)
	if !exist {
		newMikanData, err := c.mikan.ParseBangumiByUrl(ctx, targetUrl)
		if err != nil {
			return err
		}
		c.cache.Add(ParserMikan, targetUrl, newMikanData)
		mikanData = newMikanData
	} else {
		mikanData = data.(*mikan.Bangumi)
	}
	rawData.Title = mikanData.Title
	for _, episode := range mikanData.Episodes {
		rawData.Files = append(rawData.Files, episode.Title)
	}
	return nil
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
