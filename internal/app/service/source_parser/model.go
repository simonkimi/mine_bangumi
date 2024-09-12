package source_parser

type parserRawData struct {
	Title string
	Files []string
}

type ParserResult struct {
	RawTitle string
	Files    []string
	Season   int
}

type ScraperResult struct {
	ScraperTitle string
}
