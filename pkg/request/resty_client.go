package request

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type RestyClient struct {
	client *resty.Client
}

func (r *RestyClient) SetBaseURL(s string) Client {
	r.client.SetBaseURL(s)
	return r
}

func NewRestyClient(client *resty.Client) Client {
	return &RestyClient{
		client: client,
	}
}

func Default() Client {
	return &RestyClient{
		client: resty.New(),
	}
}

type RestyRequest struct {
	request *resty.Request
}

func (r *RestyRequest) Url() string {
	return r.request.URL
}

func (r *RestyRequest) SetContext(ctx context.Context) Request {
	r.request.SetContext(ctx)
	return r
}

func (r *RestyRequest) SetHeaders(m map[string]string) Request {
	r.request.SetHeaders(m)
	return r
}

func (r *RestyRequest) SetQueryParams(m map[string]string) Request {
	r.request.SetQueryParams(m)
	return r
}

func (r *RestyRequest) SetFormData(m map[string]string) Request {
	r.request.SetFormData(m)
	return r
}

func (r *RestyRequest) SetResult(a any) Request {
	r.request.SetResult(a)
	return r
}

func (r *RestyRequest) Get(s string) Response {
	rsp, err := r.request.Get(s)
	if err != nil {
		return &RestyResponse{
			request: r,
			err:     err,
		}
	}
	return &RestyResponse{
		request:  r,
		response: rsp,
	}
}

func (r *RestyRequest) Post(s string) Response {
	rsp, err := r.request.Post(s)
	if err != nil {
		return &RestyResponse{
			request: r,
			err:     err,
		}
	}
	return &RestyResponse{
		request:  r,
		response: rsp,
	}
}

func (r *RestyClient) R() Request {
	return &RestyRequest{
		request: r.client.R(),
	}
}

type RestyResponse struct {
	request  Request
	response *resty.Response
	err      error
}

func (r *RestyResponse) Request() Request {
	return r.request
}

func (r *RestyResponse) Body() []byte {
	return r.response.Body()
}

func (r *RestyResponse) String() string {
	return r.response.String()
}

func (r *RestyResponse) Error() error {
	if r.err != nil {
		return r.err
	}
	if r.response.IsError() {
		return &StatusError{Code: r.response.StatusCode()}
	}
	return nil
}
