package mikan

import (
	"context"
	"github.com/simonkimi/minebangumi/pkg/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseBangumi(t *testing.T) {
	client, err := request.NewMockFileClient("testdata/atri_bangumi.xml")
	require.Nil(t, err)

	result, err := ParseUrl(context.Background(), client, "")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "亚托莉 -我挚爱的时光-", result.Title)
}

func TestParserMyBangumi(t *testing.T) {
	client, err := request.NewMockFileClient("testdata/MyBangumi.xml")
	require.Nil(t, err)

	result, err := ParseUrl(context.Background(), client, "")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "我的番组", result.Title)
	assert.Equal(t, 9, len(result.Episodes))
	assert.Equal(t, "[SweetSub][鹿乃子大摇大摆虎视眈眈][Shikanoko Nokonoko Koshitantan][11][WebRip][1080P][AVC 8bit][简日双语]", result.Episodes[0].Title)
	t.Log(result)
}
