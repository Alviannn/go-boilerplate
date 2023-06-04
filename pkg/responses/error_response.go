package responses

import (
	"github.com/labstack/echo/v4"
	"github.com/ztrue/tracerr"
)

type (
	ErrorResponse struct {
		Message      string   `json:"message"`
		ErrorMessage string   `json:"errorMessage"`
		Stack        []string `json:"stack"`
	}

	ErrorResponseBuilder struct {
		Error CustomError
	}
)

func NewError() *CustomError {
	return &CustomError{}
}

func FromPrimitiveError(err error) *CustomError {
	if customError, ok := err.(*CustomError); ok {
		return customError
	}

	return NewError().WithError(err)
}

func sendErrorResponse(c echo.Context, err *CustomError) error {
	err.Sanitize()

	rawStackTrace := tracerr.StackTrace(err.Err)
	stackTrace := parseStackTrace(rawStackTrace)

	response := ErrorResponse{
		Message:      err.Message,
		ErrorMessage: err.Error(),
		Stack:        stackTrace,
	}

	return c.JSON(err.Code, response)
}

func parseStackTrace(rawStackList []tracerr.Frame) []string {
	var stackList []string

	for _, rawStack := range rawStackList {
		stackList = append(stackList, rawStack.String())
	}

	return stackList
}
