package handler

import "github.com/go-resty/resty/v2"

type HttpHandler struct {
	getTmpClient func() *resty.Client
}
