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

	var errorMessage string
	if err.Err != nil {
		errorMessage = err.Err.Error()
	}

	response := ErrorResponse{
		Message:      err.Message,
		ErrorMessage: errorMessage,
		Stack:        stackTrace,
	}

	return c.JSON(err.Code, response)
}

func parseStackTrace(rawStackList []tracerr.Frame) []string {
	stackList := make([]string, 0)

	for _, rawStack := range rawStackList {
		stackList = append(stackList, rawStack.String())
	}

	return stackList
}
