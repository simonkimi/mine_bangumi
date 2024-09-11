package source_parser

type parserRawData struct {
	Title string
	Files []string
}

type ParserResult struct {
	RawTitle     string
	Files        []string
	Season       int
	StableSeason bool
}

type ScraperResult struct {
	ScraperTitle string
}
