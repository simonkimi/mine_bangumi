package api

import (
	"github.com/gin-gonic/gin"
	"github.com/simonkimi/minebangumi/pkg/errno"
	"net/http"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

func (r *Response[T]) IsError() bool {
	return r.Code != errno.Success
}

func OkResponse[T any](context *gin.Context, response T) {
	context.JSON(http.StatusOK, &Response[T]{
		Code:    errno.Success,
		Data:    response,
		Message: "",
	})
}
