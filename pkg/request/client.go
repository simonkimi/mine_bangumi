package request

import (
	"context"
	"fmt"
)

type Client interface {
	SetBaseURL(string) Client
	R() Request
}

type Request interface {
	SetContext(context.Context) Request
	SetHeaders(map[string]string) Request
	SetQueryParams(map[string]string) Request
	SetFormData(map[string]string) Request
	SetResult(any) Request
	Get(string) Response
	Post(string) Response
	Url() string
}

type Response interface {
	Error() error
	Body() []byte
	String() string
	Request() Request
}

type Options struct {
	Url      string
	Context  context.Context
	Headers  map[string]string
	Query    map[string]string
	FormData map[string]string
	Result   any
}

type StatusError struct {
	Code int
}

func (s *StatusError) Error() string {
	return fmt.Sprintf("Status code error: %d", s.Code)
}
