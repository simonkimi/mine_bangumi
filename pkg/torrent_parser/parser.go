// This code references part of the implementation from AutoBangumi.
// Project URL: https://github.com/EstrellaXD/Auto_Bangumi/blob/main/backend/src/module/parser/analyser/torrent_parser.py

package torrent_parser

import (
	"github.com/simonkimi/minebangumi/tools/stringt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	TorrentFileTypeVideo = iota
	TorrentFileTypeSubtitle
)

func ParseTorrentName(torrentPath string, torrentName string, season int, fileType int) any {
	torrentPath = filepath.Base(torrentPath)
	matchNames := make([]string, 0)
	matchNames = append(matchNames, torrentPath)
	if torrentName != "" {
		matchNames = append(matchNames, torrentName)
	}

	for _, name := range matchNames {
		for _, rule := range torrentNameMatcher {
			matchGroup, err := rule.FindStringMatch(name)
			if matchGroup == nil || err != nil {
				continue
			}
			group, title := getGroup(matchGroup.GroupByNumber(1).String())
			if season == 0 {
				title, season = getSeasonAndTitle(title)
			} else {
				title, _ = getSeasonAndTitle(title)
			}
			episode, _ := strconv.Atoi(matchGroup.GroupByNumber(2).String())
			ext := filepath.Ext(torrentPath)
			if fileType == TorrentFileTypeVideo {
				return &TorrentEpisodeFile{
					Path:    torrentPath,
					Group:   group,
					Title:   title,
					Season:  season,
					Episode: episode,
					Ext:     ext,
				}
			}
			if fileType == TorrentFileTypeSubtitle {
				return &TorrentSubtitleFile{
					Path:     torrentPath,
					Group:    group,
					Title:    title,
					Season:   season,
					Episode:  episode,
					Ext:      ext,
					Language: getSubtitleLanguage(torrentPath),
				}
			}
		}
	}
	return nil
}

func getGroup(groupAndTitle string) (group string, title string) {
	n := bracketMatcher.Split(groupAndTitle, -1)
	n = stringt.RemoveEmpty(n)

	if len(n) > 1 {
		if matched, _ := regexp.MatchString(`^\d+`, n[1]); matched {
			return "", groupAndTitle
		}
		return n[0], n[1]
	}

	return "", n[0]
}

func getSeasonAndTitle(seasonAndTitle string) (title string, season int) {
	match := seasonMatcher.FindStringSubmatch(seasonAndTitle)
	season = 1
	if len(match) > 2 {
		parsedSeason, err := strconv.Atoi(match[2])
		if err != nil {
			season = 1
		} else {
			season = parsedSeason
		}
	}
	title = seasonMatcher.ReplaceAllString(seasonAndTitle, "")
	title = strings.TrimSpace(title)
	return title, season
}

func getSubtitleLanguage(subtitleName string) string {
	for lang, langList := range subtitleLang {
		for _, langStr := range langList {
			if strings.Contains(strings.ToLower(subtitleName), langStr) {
				return lang
			}
		}
	}
	return ""
}
