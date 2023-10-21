package response

import (
	"go-boilerplate/pkg/customerror"
	"net/http"
	"reflect"
)

type Builder struct {
	Data        any
	Error       error
	SuccessCode int
}

type Response struct {
	Data       any
	StatusCode int
}

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

func sanitizeError(err error) *customerror.Error {
	if err == nil {
		return nil
	}
	if customError, ok := err.(*customerror.Error); ok {
		return customError
	}

	return customerror.New().WithSourceError(err)
}

func (b *Builder) WithError(err error) *Builder {
	b.Error = sanitizeError(err)
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
func (b *Builder) sanitizeData() any {
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

func (b *Builder) Build() (res Response) {
	res.Data = b.sanitizeData()
	res.StatusCode = b.SuccessCode

	if customErr := sanitizeError(b.Error); customErr != nil {
		res.Data = customErr.ToJSON()
		res.StatusCode = customErr.Code
	}

	return
}
