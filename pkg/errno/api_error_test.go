package errno

import (
	"errors"
	errors2 "github.com/cockroachdb/errors"
	"github.com/simonkimi/minebangumi/pkg/logger"
	"log"
	"testing"
)

func init() {
	logger.Setup()
}

func TestNewApiErrorWithCausef(t *testing.T) {
	existErr := errors2.New("exist error")
	apiError := NewApiErrorWithCausef(ErrorApiParse, existErr, "test error")
	log.Printf(ErrMsg(apiError))
}

func TestNewApiErrorWithCause(t *testing.T) {
	existErr := errors.New("exist error")
	apiError := NewApiErrorWithCause(ErrorApiParse, existErr)
	log.Printf(ErrMsg(apiError))
}

func TestNewApiError(t *testing.T) {
	apiError := NewApiError(ErrorApiParse)
	log.Printf(ErrMsg(apiError))
}
