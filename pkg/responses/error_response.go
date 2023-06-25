package responses

import (
	"net/http"

	"github.com/ztrue/tracerr"
)

type (
	ErrorResponse struct {
		Message            string   `json:"message"`
		SourceErrorMessage string   `json:"sourceErrorMessage"`
		Stack              []string `json:"stack"`
	}
)

func NewError() *CustomError {
	customError := &CustomError{
		Message: "Unhandled error",
		Code:    http.StatusInternalServerError,
	}
	customError.ThisError = tracerr.Wrap(customError)

	return customError
}

func ParseAsCustomError(err error) *CustomError {
	if customError, ok := err.(*CustomError); ok {
		return customError
	}

	return NewError().WithSourceError(err)
}

func BuildErrorResponse(err *CustomError) ErrorResponse {
	return ErrorResponse{
		Message:            err.Message,
		SourceErrorMessage: err.GetWorkingError().Error(),
		Stack:              err.GetStackTrace(),
	}
}
