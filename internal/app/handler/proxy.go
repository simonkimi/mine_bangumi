package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/internal/app/api"
	"github.com/simonkimi/minebangumi/internal/app/service"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"net/http"
	"path/filepath"
)

func Poster(c *gin.Context) {
	ctx := c.Request.Context()

	var query api.PosterQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		_ = c.Error(errno.NewApiError(errno.BadRequest))
		return
	}

	data, err := service.GetPoster(ctx, query.TargetType, query.Target)
	if err != nil {
		_ = c.Error(err)
		return
	}
	mime, err := getMime(query.Target)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, mime, data)
}

func getMime(path string) (string, error) {
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg", nil
	case ".png":
		return "image/png", nil
	case ".gif":
		return "image/gif", nil
	case ".webp":
		return "image/webp", nil
	case ".bmp":
		return "image/bmp", nil
	case ".svg":
		return "image/svg+xml", nil
	case ".jp2", ".j2k", ".jpf", ".jpx", ".jpm", ".mj2":
		return "image/jp2", nil
	default:
		return "", errno.NewApiErrorf(errno.BadRequest, "unsupported image type: %s", ext)
	}
}
