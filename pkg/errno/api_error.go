package errno

import (
	"fmt"
	"github.com/cockroachdb/errors"
)

type ApiError struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	internalError error
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func (e *ApiError) Unwrap() error {
	return e.internalError
}

func (e *ApiError) internalErrMsg() string {
	return fmt.Sprintf("code: %d, message: %s, error: %+v", e.Code, e.Message, e.internalError)
}

func ErrMsg(err error) string {
	if err == nil {
		return ""
	}
	var e *ApiError
	if errors.As(err, &e) {
		return e.internalErrMsg()
	}
	return err.Error()
}

func NewApiError(code int) *ApiError {
	errMsg := getErrorMessage(code)
	return &ApiError{
		Code:          code,
		Message:       errMsg,
		internalError: errors.NewWithDepth(1, errMsg),
	}
}

func NewApiErrorf(code int, format string, args ...any) *ApiError {
	return &ApiError{
		Code:          code,
		Message:       getErrorMessage(code),
		internalError: errors.NewWithDepthf(1, format, args...),
	}
}

func NewApiErrorWithCause(code int, cause error) *ApiError {
	errMsg := getErrorMessage(code)
	return &ApiError{
		Code:          code,
		Message:       errMsg,
		internalError: errors.WrapWithDepth(1, cause, errMsg),
	}
}

func NewApiErrorWithCausef(code int, cause error, format string, args ...any) *ApiError {
	return &ApiError{
		Code:          code,
		Message:       getErrorMessage(code),
		internalError: errors.WrapWithDepthf(1, cause, format, args...),
	}
}
