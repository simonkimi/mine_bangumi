package testutil

import (
	"bytes"
	"compress/gzip"
	"github.com/gavv/httpexpect/v2"
	"io"
	"net/http/httptest"
	"testing"
)

func DecompressResponse(w *httptest.ResponseRecorder) ([]byte, error) {
	if w.Result().Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(bytes.NewReader(w.Body.Bytes()))
		if err != nil {
			return nil, err
		}
		defer func(gzipReader *gzip.Reader) {
			_ = gzipReader.Close()
		}(gzipReader)
		decompressedBody, err := io.ReadAll(gzipReader)
		if err != nil {
			return nil, err
		}
		return decompressedBody, nil
	}
	return w.Body.Bytes(), nil
}

func NewDebugHttp(t *testing.T, url string) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		TestName: t.Name(),
		BaseURL:  url,
		Printers: []httpexpect.Printer{httpexpect.NewDebugPrinter(t, true)},
		Reporter: httpexpect.NewAssertReporter(t),
	})
}
