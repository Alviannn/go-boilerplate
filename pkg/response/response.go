package response

import (
	"go-boilerplate/pkg/customerror"
	"net/http"
	"reflect"
)

var (
	emptySlice = make([]string, 0)
	emptyMap   = make(map[string]string)
)

type (
	Builder struct {
		Data        any
		Error       *customerror.Error
		SuccessCode int
	}

	Response struct {
		Data       any
		StatusCode int
	}
)

func NewBuilder() *Builder {
	return &Builder{
		SuccessCode: http.StatusOK,
	}
}

func (b *Builder) WithData(data any) *Builder {
	b.Data = data
	return b
}

func (b *Builder) WithSuccessCode(statusCode int) *Builder {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	b.SuccessCode = statusCode
	return b
}

func (b *Builder) WithError(err error) *Builder {
	if err == nil {
		return b
	}
	if customErr, ok := err.(*customerror.Error); ok {
		b.Error = customErr
	} else {
		b.Error = customerror.New().WithSourceError(err)
	}
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
func (b *Builder) sanitizeData(data any) any {
	if data == nil {
		return emptyMap
	}

	dataRef := reflect.ValueOf(data)
	if !dataRef.IsZero() {
		return data
	}

	switch dataRef.Kind() {
	case reflect.Struct:
		return emptyMap
	case reflect.Slice, reflect.Array:
		return emptySlice
	}
	return data
}

func (b *Builder) Build() (res Response) {
	res.Data = b.sanitizeData(b.Data)
	res.StatusCode = b.SuccessCode

	if err := b.Error; err != nil {
		res.Data = err.ToJSON()
		res.StatusCode = err.Code
	}
	return
}
