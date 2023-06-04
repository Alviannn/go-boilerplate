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
	return &ResponseBuilder{}
}

func (b *ResponseBuilder) WithData(data any) *ResponseBuilder {
	b.Data = data
	return b
}

func (b *ResponseBuilder) WithSuccessCode(statusCode int) *ResponseBuilder {
	b.SuccessCode = statusCode
	return b
}

func (b *ResponseBuilder) WithError(err error) *ResponseBuilder {
	b.Error = err
	return b
}

func (b *ResponseBuilder) Send(c echo.Context) error {
	if b.Error != nil {
		return sendErrorResponse(c, FromPrimitiveError(b.Error))
	}

	dataReflect := reflect.ValueOf(b.Data)
	dataKind := dataReflect.Kind()

	if b.SuccessCode == 0 {
		b.WithSuccessCode(http.StatusOK)
	}

	var sanitizedData any = b.Data
	if (dataKind == reflect.Struct || dataKind == reflect.Slice) && dataReflect.IsZero() {
		if dataKind == reflect.Struct {
			sanitizedData = make(map[any]any, 0)
		} else {
			sanitizedData = make([]any, 0)
		}
	} else if b.Data == nil {
		sanitizedData = make(map[any]any, 0)
	}

	return c.JSON(b.SuccessCode, sanitizedData)
}
