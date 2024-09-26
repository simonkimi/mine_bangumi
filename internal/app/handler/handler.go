package handler

import "github.com/go-resty/resty/v2"

type HttpHandler struct {
	getTmpClient func() *resty.Client
}

func NewHttpHandler(getTmpClient func() *resty.Client) *HttpHandler {
	return &HttpHandler{getTmpClient}
}
