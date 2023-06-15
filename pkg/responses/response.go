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
	if b.SuccessCode == 0 {
		b.WithSuccessCode(http.StatusOK)
	}

	dataReflect := reflect.ValueOf(b.Data)
	dataKind := dataReflect.Kind()

	emptySlice := make([]string, 0)
	emptyMap := make(map[string]string)

	var sanitizedData any = b.Data
	if (dataKind == reflect.Struct || dataKind == reflect.Slice) && dataReflect.IsZero() {
		if dataKind == reflect.Struct {
			sanitizedData = emptyMap
		} else {
			sanitizedData = emptySlice
		}
	} else if b.Data == nil {
		sanitizedData = emptyMap
	}

	return c.JSON(b.SuccessCode, sanitizedData)
}
