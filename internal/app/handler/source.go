package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/pkg/errno"
)

// Source godoc
// @Summary Parse source
// @Description Parse the source using the specified parser
// @Tags parser
// @Accept json
// @Produce json
// @Param form body api.ParseSourceForm true "Parse Source Form"
// @Success 200 {object} api.ParseSourceResponse
// @Router /api/v1/source/parse [post]
func Source(c *gin.Context) {
	ctx := c.Request.Context()
	var form *api.ParseSourceForm
	if err := c.ShouldBindJSON(&form); err != nil {
		_ = c.Error(errno.NewFormError(err))
		return
	}
	source, err := service.ParseSource(ctx, form.Source, form.Parser)
	if err != nil {
		_ = c.Error(err)
		return
	}
	api.OkResponse(c, source)
}

// Scrape godoc
// @Summary Scrape source
// @Description Scrape the source using the specified scraper
// @Tags scraper
// @Accept json
// @Produce json
// @Param form body api.ScrapeForm true "Scrape Form"
// @Success 200 {object} api.ScrapeResponse
// @Router /api/v1/source/scrape [post]
func Scrape(c *gin.Context) {
	ctx := c.Request.Context()
	var form *api.ScrapeForm
	if err := c.ShouldBindJSON(&form); err != nil {
		_ = c.Error(errno.NewFormError(err))
		return
	}
	info, err := service.ScrapeService(ctx, form)
	if err != nil {
		_ = c.Error(err)
		return
	}
	api.OkResponse(c, info)
}
