package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response[T any] struct {
	Code      APIStatusEnum  `json:"code"`
	Data      T              `json:"data"`
	Message   string         `json:"message"`
	Extension map[string]any `json:"extension"`
}

func (r *Response[T]) IsError() bool {
	return r.Code != APIStatusEnumSuccess
}

func OkResponse[T any](context *gin.Context, response T) {
	context.JSON(http.StatusOK, &Response[T]{
		Code:    APIStatusEnumSuccess,
		Data:    response,
		Message: "",
	})
}

func OkResponseNil(c *gin.Context) {
	c.JSON(http.StatusOK, &Response[any]{
		Code:    APIStatusEnumSuccess,
		Data:    nil,
		Message: "",
	})
}
