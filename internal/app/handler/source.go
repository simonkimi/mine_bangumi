package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/service/source_parser"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"net/http"
)

// ParseSource godoc
// @Summary Parse source
// @Description Parse the source using the specified parser
// @Tags parser
// @Accept json
// @Produce json
// @Param form body dto.ParseSourceForm true "Parse Source Form"
// @Success 200 {object} dto.ParseSourceResponse
// @Router /parse [post]
func ParseSource(c *gin.Context) {
	ctx := c.Request.Context()
	var form *api.ParseSourceForm
	if err := c.ShouldBindJSON(&form); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errno.NewFormError(err))
		return
	}
	source, err := source_parser.ParseSource(ctx, form.Source, form.Parser)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &api.ParseSourceResponse{
		Title:  source.RawTitle,
		Files:  source.Files,
		Season: source.Season,
	})
}
