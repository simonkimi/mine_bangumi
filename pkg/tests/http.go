package tests

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http/httptest"
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
