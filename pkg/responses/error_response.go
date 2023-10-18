package responses

import (
	"net/http"

	"github.com/ztrue/tracerr"
)

type (
	ErrorResponse struct {
		Message            string   `json:"message"`
		SourceErrorMessage string   `json:"source_error_message"`
		Stack              []string `json:"stack"`
	}
)

func NewError() *CustomError {
	customError := &CustomError{
		Message: "Unhandled error",
		Code:    http.StatusInternalServerError,
	}
	customError.thisError = tracerr.Wrap(customError)

	return customError
}

func parseAsCustomErrorOrNil(err error) *CustomError {
	if err == nil {
		return nil
	}
	if customError, ok := err.(*CustomError); ok {
		return customError
	}

	return NewError().WithSourceError(err)
}

func (res ErrorResponse) IsValid() bool {
	return res.Message != "" || len(res.Stack) > 0 || res.SourceErrorMessage != ""
}
