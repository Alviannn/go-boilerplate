package responses

import (
	"net/http"

	"github.com/ztrue/tracerr"
)

type CustomError struct {
	Err     error
	Message string
	Code    int
}

func (e *CustomError) WithMessage(message string) *CustomError {
	e.Message = message
	return e
}

func (e *CustomError) WithCode(statusCode int) *CustomError {
	e.Code = statusCode
	return e
}

func (e *CustomError) WithError(err error) *CustomError {
	if customError, ok := err.(*CustomError); ok {
		// Reuse the custom error instead of recreating it
		// in order to keep the stack trace.
		err = customError.Err
	} else {
		err = tracerr.Wrap(err)
	}

	e.Err = err
	return e
}

func (e *CustomError) Sanitize() *CustomError {
	if e.Code == 0 {
		e.WithCode(http.StatusInternalServerError)
	}
	if e.Message == "" {
		e.WithMessage("Unhandled server error")
	}

	return e
}

func (e CustomError) Error() string {
	return e.Message
}
