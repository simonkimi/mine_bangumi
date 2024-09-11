package source_parser

import (
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

func ParseRss(targetUrl string, parser string) error {
	if stringt.IsEmptyOrWhitespace(parser) {
		parser = guestParser(targetUrl)
	}
	switch parser {
	case ParserMikan:
		//return mikan.ParseFromUrl(targetUrl)
	case ParserMagnet:
		//return magnet.ParseFromUrl(targetUrl)
	}

	return nil
}

//https://mikanani.me/RSS/Bangumi?bangumiId=3373&subgroupid=570
