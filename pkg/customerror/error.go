package customerror

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ztrue/tracerr"
)

// Error is the standard error for this `responses` package,
// where we can give as much details as possible in the response.
//
// It's created using `responses.NewError()` function since it has
// default values for this struct instance. The struct fields provides
// the necessary details for the error response, such as the errors
// being traceable (generate stack trace).
//
// This struct is using builder pattern, you can see the code example
// in `responses.NewError()` function.
type Error struct {
	// Context is the context of the error.
	Context context.Context

	// SourceError is the real source of error.
	//
	// It acts as a detail that shows where the error really came from,
	// then generates the stack trace using `tracerr`. This field can be
	// `nil`, read `ThisError` on why.
	SourceError error

	// thisError is the current `CustomError` instance but in a form
	// of a traceable error (basically with stack trace).
	//
	// It is responsible for replacing `SourceError` when it's `nil`.
	// This will never be (and should never be) `nil`. The field setup
	// exists in the struct creation function.
	thisError error

	// Message is human-readable or human-friendly error that is made
	// specifically for someone else to read with ease.
	//
	// An example of this is "User's email is already registered".
	// It defauls to "Unhandled error" to fit with the default value
	// of HTTP status code in `Code`.
	Message string

	// Code is the HTTP status code response when the error happens.
	//
	// It defaults to `http.StatusInternalServerError` or 500.
	Code int

	// IsPanic is a flag that indicates whether the error is a panic.
	IsPanic bool
}

func EqualCode(err error, code int) bool {
	if err == nil {
		return false
	}
	if customError, ok := err.(*Error); ok {
		return customError.Code == code
	}
	return false
}

func New() *Error {
	customError := &Error{
		Message: "Unhandled error",
		Code:    http.StatusInternalServerError,
		IsPanic: false,
	}
	customError.thisError = tracerr.Wrap(customError)

	return customError
}

func (e *Error) WithMessage(message string) *Error {
	e.Message = message
	return e
}

func (e *Error) WithCode(statusCode int) *Error {
	// Defaults to internal server error (500) when it's empty
	// because we're sending it as HTTP response.
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}

	e.Code = statusCode
	return e
}

func (e *Error) WithContext(ctx context.Context) *Error {
	e.Context = ctx
	return e
}

func (e *Error) WithSourceError(err error) *Error {
	if customError, ok := err.(*Error); ok {
		// Reuse the traceable errors instead of recreating it
		// in order to keep the stack trace.
		err = customError.GetWorkingError()
	} else if _, ok := err.(tracerr.Error); !ok {
		// Make it traceable error when it's not already to
		// generate stack trace.
		err = tracerr.Wrap(err)
	}

	e.SourceError = err
	return e
}

func (e *Error) WithPanic(isPanic bool) *Error {
	e.IsPanic = isPanic
	return e
}

// GetWorkingError gets the error instance that can be used or
// the one that works between the source and current error.
//
// The first priority is `SourceError`, although when the value
// is `nil` it will be replaced with `ThisError`.
func (e Error) GetWorkingError() (err error) {
	err = e.SourceError
	if err == nil {
		err = e.thisError
	}

	return err
}

func (e Error) GetStackTrace() []string {
	// Use `make` to avoid returning `nil` value
	stackList := make([]string, 0)
	rawStackList := tracerr.StackTrace(e.GetWorkingError())

	for _, rawStack := range rawStackList {
		stackList = append(stackList, rawStack.String())
	}

	return stackList
}

func (e Error) GetStackLine() string {
	rawStackList := tracerr.StackTrace(e.GetWorkingError())
	if len(rawStackList) < 2 {
		return ""
	}

	rawStack := rawStackList[1]
	return fmt.Sprintf("%s at line %d", rawStack.Func, rawStack.Line)
}

func (e Error) Error() string {
	return e.Message
}

func (e *Error) ToJSON() ErrorJSON {
	return ErrorJSON{
		src:                e,
		Message:            e.Message,
		SourceErrorMessage: e.GetWorkingError().Error(),
		StackLine:          e.GetStackLine(),
		Stack:              e.GetStackTrace(),
	}
}
