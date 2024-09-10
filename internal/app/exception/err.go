package exception

import "fmt"
import "github.com/cockroachdb/errors"

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

func NewApiError(code int) *ApiError {
	stackTrace := errors.NewWithDepth(1, "ApiError occurred")
	return &ApiError{
		Code:          code,
		Message:       getErrorMessage(code),
		internalError: stackTrace,
	}
}

func NewApiErrorWithCause(code int, cause error) *ApiError {
	return &ApiError{
		Code:          code,
		Message:       getErrorMessage(code),
		internalError: errors.Wrap(cause, "ApiError with cause"),
	}
}
