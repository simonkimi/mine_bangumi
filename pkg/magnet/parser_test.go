package magnet

import (
	"context"
	"testing"
)

func TestParseMagnet(t *testing.T) {
	url := `magnet:?xt=urn:btih:ed3f901626a717d897118bf08d3b476a23e18d31&tr=http%3a%2f%2ft.nyaatracker.com%2fannounce&tr=http%3a%2f%2ftracker.kamigami.org%3a2710%2fannounce&tr=http%3a%2f%2fshare.camoe.cn%3a8080%2fannounce&tr=http%3a%2f%2fopentracker.acgnx.se%2fannounce&tr=http%3a%2f%2fanidex.moe%3a6969%2fannounce&tr=http%3a%2f%2ft.acg.rip%3a6699%2fannounce&tr=https%3a%2f%2ftr.bangumi.moe%3a9696%2fannounce&tr=udp%3a%2f%2ftr.bangumi.moe%3a6969%2fannounce&tr=http%3a%2f%2fopen.acgtracker.com%3a1096%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce`
	fileInfo, err := ParseMagnet(context.Background(), url)
	if err != nil {
		t.Errorf("ParseMagnet() error = %v", err)
		return
	}
	t.Logf("ParseMagnet() = %v", fileInfo)
}
