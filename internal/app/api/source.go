package api

type ParseSourceForm struct {
	Source string `json:"source" binding:"required"`
	Parser string `json:"parser" binding:"required"`
}

type ParseSourceResponse struct {
	Title  string   `json:"title"`
	Files  []string `json:"files"`
	Season int      `json:"season"`
}

type ScrapeForm struct {
	Title    string `json:"title" binding:"required"`
	Scraper  string `json:"scraper" binding:"required"`
	Language string `json:"language"`
}

type ScrapeResponse struct {
	Scraper       string                  `json:"scraper"`
	Title         string                  `json:"title"`
	OriginalTitle string                  `json:"original_title"`
	FirstAirYear  string                  `json:"first_air_year"`
	Overview      string                  `json:"overview"`
	PosterPath    string                  `json:"poster_path"`
	Seasons       []*ScrapeSeasonResponse `json:"seasons"`
}

type ScrapeSeasonResponse struct {
	SeasonId int    `json:"season_id"`
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Poster   string `json:"poster"`
}
