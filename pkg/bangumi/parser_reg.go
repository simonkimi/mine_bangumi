package bangumi

import (
	"github.com/dlclark/regexp2"
	"regexp"
)

var torrentNameMatcher = []*regexp2.Regexp{
	regexp2.MustCompile(`(.*) - (\d{1,4}(?!\d|p)|\d{1,4}\.\d{1,2}(?!\d|p))(?:v\d{1,2})?(?: )?(?:END)?(.*)`, regexp2.IgnoreCase),
	regexp2.MustCompile(`(.*)[\[\ E](\d{1,4}|\d{1,4}\.\d{1,2})(?:v\d{1,2})?(?: )?(?:END)?[\]\ ](.*)`, regexp2.IgnoreCase),
	regexp2.MustCompile(`(.*)\[(?:第)?(\d*\.*\d*)[话集話](?:END)?\](.*)`, regexp2.IgnoreCase),
	regexp2.MustCompile(`(.*)第?(\d*\.*\d*)[话話集](?:END)?(.*)`, regexp2.IgnoreCase),
	regexp2.MustCompile(`(.*)(?:S\d{2})?EP?(\d+)(.*)`, regexp2.IgnoreCase),
}

var bracketMatcher = regexp.MustCompile(`[\[\]()【】（）]`)

var intMatcher = regexp.MustCompile(`\d+`)

var seasonMatcher = regexp.MustCompile(`(?i)([Ss]|Season )(\d{1,3})`)

var subtitleLang = map[string][]string{
	"zh-Hant": {"tc", "cht", "繁", "zh-tw"},
	"zh-Hans": {"sc", "chs", "简", "zh"},
}
