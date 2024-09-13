package tmdb

import (
	"github.com/simonkimi/minebangumi/domain"
	"testing"
)

func TestSearch(t *testing.T) {
	result, err := Search("魔法禁书目录")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestDetail(t *testing.T) {
	result, err := QueryForDetail(30980, domain.LanguageZhHans)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
