package api

import (
	"context"
	"fmt"
)

type Error struct {
	Message       string
	Code          APIStatusEnum
	internalError error
	Extensions    map[string]any
}

func (g *Error) Error() string {
	return g.Message
}

func NewCancelError() error {
	return &Error{
		Message:       "Request canceled",
		Code:          APIStatusEnumCancel,
		internalError: context.Canceled,
	}
}

func NewCancelErrorf(format string, args ...any) error {
	return &Error{
		Message:       fmt.Sprintf(format, args...),
		Code:          APIStatusEnumCancel,
		internalError: context.Canceled,
	}
}

func NewTimeoutErrorf(format string, args ...any) error {
	return &Error{
		Message:       fmt.Sprintf(format, args...),
		Code:          APIStatusEnumTimeout,
		internalError: context.DeadlineExceeded,
	}
}

func NewThirdPartyErrorf(err error, url string, format string, args ...any) error {
	return &Error{
		Message:       fmt.Sprintf(format, args...),
		Code:          APIStatusEnumThirdPartyAPIError,
		internalError: err,
		Extensions: map[string]any{
			"url": url,
		},
	}
}

func NewBadRequestErrorf(format string, args ...any) error {
	return &Error{
		Message: fmt.Sprintf(format, args...),
		Code:    APIStatusEnumBadRequest,
	}
}

func NewFormValidationError(errMsg map[string]string) error {
	return &Error{
		Message:       "Form validation error",
		Code:          APIStatusEnumFormValidationError,
		internalError: nil,
		Extensions: map[string]any{
			"fields": errMsg,
		},
	}
}

func NewUnAuthError() error {
	return &Error{
		Message:    "Unauthorized",
		Code:       APIStatusEnumUnauthorized,
		Extensions: nil,
	}
}

func NewForbiddenError() error {
	return &Error{
		Message:    "Forbidden",
		Code:       APIStatusEnumForbidden,
		Extensions: nil,
	}
}

func NewInternalServerError(err error) error {
	return &Error{
		Message:       "Internal server error",
		Code:          APIStatusEnumInternalServerError,
		internalError: err,
	}
}
