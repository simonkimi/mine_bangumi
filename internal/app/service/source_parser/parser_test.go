package source_parser

import (
	"context"
	"github.com/simonkimi/minebangumi/internal/app/config"
	"github.com/simonkimi/minebangumi/pkg/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	tests.LoadTestEnv()
}

func TestParseMikanUrl(t *testing.T) {
	tests.WorkOnTempDir(t, true)
	config.Setup()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/atri_bangumi.xml")
	}))
	parseResult, err := ParseSource(context.Background(), server.URL, ParserMikan)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parseResult)
}

func TestParseMagUrl(t *testing.T) {
	url := "magnet:?xt=urn:btih:3d08fc350ef88251620f3ea02d6883831daf9ac4&tr=http%3a%2f%2ft.nyaatracker.com%2fannounce&tr=http%3a%2f%2ftracker.kamigami.org%3a2710%2fannounce&tr=http%3a%2f%2fshare.camoe.cn%3a8080%2fannounce&tr=http%3a%2f%2fopentracker.acgnx.se%2fannounce&tr=http%3a%2f%2fanidex.moe%3a6969%2fannounce&tr=http%3a%2f%2ft.acg.rip%3a6699%2fannounce&tr=https%3a%2f%2ftr.bangumi.moe%3a9696%2fannounce&tr=udp%3a%2f%2ftr.bangumi.moe%3a6969%2fannounce&tr=http%3a%2f%2fopen.acgtracker.com%3a1096%2fannounce&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce"
	parseResult, err := ParseSource(context.Background(), url, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(parseResult)
}
