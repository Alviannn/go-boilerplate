package responses

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

type ResponseBuilder struct {
	Data        any
	Error       error
	SuccessCode int
}

func New() *ResponseBuilder {
	return &ResponseBuilder{
		SuccessCode: http.StatusOK,
	}
}

func (b *ResponseBuilder) WithData(data any) *ResponseBuilder {
	b.Data = data
	return b
}

func (b *ResponseBuilder) WithSuccessCode(statusCode int) *ResponseBuilder {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	b.SuccessCode = statusCode
	return b
}

func (b *ResponseBuilder) WithError(err error) *ResponseBuilder {
	b.Error = parseAsCustomErrorOrNil(err)
	return b
}

// sanitizeData will sanitize the `Data` for JSON response.
//
// By sanitizing, it means to make the value of `Data` not `nil`.
// If it's an empty slice, it will change to `[]`.
// If it's an empty object/struct instance, it will change to `{}`.
//
// Anything that's not a slice or struct or `nil` value
// will become the default zero value of the type,
// ex: empty string stays as "".
func (b *ResponseBuilder) sanitizeData() any {
	dataReflect := reflect.ValueOf(b.Data)
	dataKind := dataReflect.Kind()

	emptySlice := make([]string, 0)
	emptyMap := make(map[string]string)

	var newData any = b.Data
	isDataStructOrSlice := (dataKind == reflect.Struct || dataKind == reflect.Slice)

	if newData == nil {
		newData = emptyMap
	} else if isDataStructOrSlice && dataReflect.IsZero() {
		if dataKind == reflect.Struct {
			newData = emptyMap
		} else {
			newData = emptySlice
		}
	}

	return newData
}

func (b *ResponseBuilder) Send(c echo.Context) error {
	if customErr := parseAsCustomErrorOrNil(b.Error); customErr != nil {
		return c.JSON(customErr.Code, customErr.BuildResponse())
	}
	return c.JSON(b.SuccessCode, b.sanitizeData())
}
