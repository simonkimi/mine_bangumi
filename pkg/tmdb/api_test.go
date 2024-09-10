package tmdb

import (
	"github.com/simonkimi/minebangumi/domain"
	"github.com/simonkimi/minebangumi/internal/pkg/setup"
	"testing"
)

func init() {
	setup.SetupTest()
}

func TestSearch(t *testing.T) {
	result, err := Search("魔法禁书目录")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestDetail(t *testing.T) {
	result, err := GetDetail(30980, domain.LanguageZhHans)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
