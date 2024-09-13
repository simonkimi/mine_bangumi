package mikan

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseBangumi(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/atri_bangumi.xml")
	}))
	defer mockServer.Close()

	result, err := ParseBangumiByUrl(context.Background(), mockServer.URL)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "亚托莉 -我挚爱的时光-", result.Title)
}

func TestParserMyBangumi(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/MyBangumi.xml")
	}))
	defer mockServer.Close()

	result, err := ParseBangumiByUrl(context.Background(), mockServer.URL)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "我的番组", result.Title)
	assert.Equal(t, 9, len(result.Episodes))
	assert.Equal(t, "[SweetSub][鹿乃子大摇大摆虎视眈眈][Shikanoko Nokonoko Koshitantan][11][WebRip][1080P][AVC 8bit][简日双语]", result.Episodes[0].Title)
	t.Log(result)
}
