package api

type PosterQuery struct {
	Target     string `form:"target" binding:"required"`
	TargetType string `form:"type" binding:"required"`
}
