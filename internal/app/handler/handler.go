package handler

import (
	"github.com/simonkimi/minebangumi/internal/app/service"
)

type HttpHandler struct {
	httpx *service.HttpX
}

func NewHttpHandler(httpx *service.HttpX) *HttpHandler {
	return &HttpHandler{
		httpx: httpx,
	}
}
