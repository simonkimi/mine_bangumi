package api

import (
	"context"
	"fmt"
)

type GqlError struct {
	Message       string
	Code          APIStatusEnum
	internalError error
	Extensions    map[string]any
}

func (g *GqlError) Error() string {
	return g.Message
}

func NewCancelError() error {
	return &GqlError{
		Message:       "Request canceled",
		Code:          APIStatusEnumCancel,
		internalError: context.Canceled,
	}
}

func NewCancelErrorf(format string, args ...any) error {
	return &GqlError{
		Message:       fmt.Sprintf(format, args...),
		Code:          APIStatusEnumCancel,
		internalError: context.Canceled,
	}
}

func NewTimeoutErrorf(format string, args ...any) error {
	return &GqlError{
		Message:       fmt.Sprintf(format, args...),
		Code:          APIStatusEnumTimeout,
		internalError: context.DeadlineExceeded,
	}
}

func NewThirdPartyErrorf(err error, url string, format string, args ...any) error {
	return &GqlError{
		Message:       fmt.Sprintf(format, args...),
		Code:          APIStatusEnumThirdPartyAPIError,
		internalError: err,
		Extensions: map[string]any{
			"url": url,
		},
	}
}

func NewBadRequestErrorf(format string, args ...any) error {
	return &GqlError{
		Message: fmt.Sprintf(format, args...),
		Code:    APIStatusEnumBadRequest,
	}
}
