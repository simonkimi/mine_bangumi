package dto

type ParseSourceForm struct {
	Source string `json:"source" binding:"required"`
	Parser string `json:"parser" binding:"required"`
}

type ParseSourceResponse struct {
	Title  string   `json:"title"`
	Files  []string `json:"files"`
	Season int      `json:"season"`
}
