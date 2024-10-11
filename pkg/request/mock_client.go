package request

import (
	"context"
	"os"
	"strings"
)

type MockClient struct {
	data []byte
}

func NewMockFileClient(path string) (Client, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &MockClient{data}, nil
}

func (m *MockClient) SetBaseURL(s string) Client {
	return m
}

func (m *MockClient) R() Request {
	return &MockRequest{
		data: m.data,
	}
}

type MockRequest struct {
	data []byte
}

type MockResponse struct {
	request *MockRequest
	data    []byte
}

func (m *MockResponse) Error() error {
	return nil
}

func (m *MockResponse) Body() []byte {
	return m.data
}

func (m *MockResponse) String() string {
	return strings.TrimSpace(string(m.data))
}

func (m *MockResponse) Request() Request {
	return m.request
}

func (m *MockRequest) SetContext(context.Context) Request {
	return m
}

func (m *MockRequest) SetHeaders(map[string]string) Request {
	return m
}

func (m *MockRequest) SetQueryParams(map[string]string) Request {
	return m
}

func (m *MockRequest) SetFormData(map[string]string) Request {
	return m
}

func (m *MockRequest) SetResult(any) Request {
	return m
}

func (m *MockRequest) Get(string) Response {
	return &MockResponse{
		request: m,
		data:    m.data,
	}
}

func (m *MockRequest) Post(string) Response {
	return &MockResponse{
		request: m,
		data:    m.data,
	}
}

func (m *MockRequest) Url() string {
	return ""
}
