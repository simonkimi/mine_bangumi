package source_parser

import (
	"context"
	"github.com/simonkimi/minebangumi/internal/pkg/cache"
	"github.com/simonkimi/minebangumi/pkg/bangumi"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"github.com/simonkimi/minebangumi/pkg/magnet"
	"github.com/simonkimi/minebangumi/pkg/mikan"
	"github.com/simonkimi/minebangumi/tools/stringt"
	"regexp"
	"strings"
)

const (
	ParserMikan  = "mikan"
	ParserMagnet = "magnet"
)

type RssModel struct {
	RssTitle string
}

func guestParser(targetUrl string) string {
	mikanReg := regexp.MustCompile("^https?://mikanani\\.(me|tv)")
	if mikanReg.MatchString(targetUrl) {
		return ParserMikan
	}
	if strings.HasPrefix(targetUrl, "magnet:") {
		return ParserMagnet
	}
	return ""
}

func ParseUrl(ctx context.Context, targetUrl string, parser string) (*ParserResult, error) {
	if stringt.IsEmptyOrWhitespace(parser) {
		parser = guestParser(targetUrl)
	}
	// 解析原始数据
	rawData := &parserRawData{}
	switch parser {
	case ParserMikan:
		err := mikanParse(ctx, targetUrl, rawData)
		if err != nil {
			return nil, err
		}
	case ParserMagnet:
		err := magnetParse(ctx, targetUrl, rawData)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errno.NewApiErrorf(errno.BadRequest, "Unknown parser: %s", parser)
	}
	// 解析原始数据, 准备刮削
	bangumiTitle := rawData.Title
	season := -1
	if len(rawData.Files) != 0 {
		bf := bangumi.ParseBangumiSourceName(rawData.Files[0], "")
		if bf != nil {
			bangumiTitle = bf.Title
			season = bf.Season
		}
	}
	stableSeason := true
	if season == -1 {
		stableSeason = false
		season = 1
	}

	// 刮削数据

	return &ParserResult{
		RawTitle:     bangumiTitle,
		Season:       season,
		StableSeason: stableSeason,
	}, nil
}

func mikanParse(ctx context.Context, targetUrl string, rawData *parserRawData) error {
	mikanData, exist := cache.Get[mikan.Bangumi](ParserMikan, targetUrl)
	if !exist {
		newMikanData, err := mikan.ParseBangumiByUrl(ctx, targetUrl)
		if err != nil {
			return err
		}
		cache.Add(ParserMikan, targetUrl, newMikanData)
		mikanData = newMikanData
	}
	rawData.Title = mikanData.Title
	for _, episode := range mikanData.Episodes {
		rawData.Files = append(rawData.Files, episode.Title)
	}
	return nil
}

func magnetParse(ctx context.Context, targetUrl string, rawData *parserRawData) error {
	magnetData, exist := cache.Get[magnet.FileInfo](ParserMagnet, targetUrl)
	if !exist {
		newMagnetData, err := magnet.ParseMagnet(ctx, targetUrl)
		if err != nil {
			return err
		}
		cache.Add(ParserMagnet, targetUrl, newMagnetData)
		magnetData = newMagnetData
	}
	rawData.Title = magnetData.Name
	rawData.Files = magnetData.Files
	return nil
}
