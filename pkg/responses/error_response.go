package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ztrue/tracerr"
)

type (
	ErrorResponse struct {
		Message            string   `json:"message"`
		SourceErrorMessage string   `json:"sourceErrorMessage"`
		Stack              []string `json:"stack"`
	}

	ErrorResponseBuilder struct {
		Error CustomError
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

func FromPrimitiveError(err error) *CustomError {
	if customError, ok := err.(*CustomError); ok {
		return customError
	}

	return NewError().WithSourceError(err)
}

func sendErrorResponse(c echo.Context, err *CustomError) error {
	response := ErrorResponse{
		Message:            err.Message,
		SourceErrorMessage: err.GetWorkingError().Error(),
		Stack:              err.GetStackTrace(),
	}

	return c.JSON(err.Code, response)
}
