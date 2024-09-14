package handler

import (
	"path/filepath"
	"testing"
)

func TestMime(t *testing.T) {
	url := "https://image.tmdb.org/t/p/w780/971wquiGNfj9xPfPZ92XVFhynPP.jpg"
	ext := filepath.Ext(url)
	t.Log(ext)
}
