package testdata

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	PersonSimple struct {
		Name string `json:"name" validate:"required,min=3,max=100"`
		Age  int    `json:"age" validate:"gte=0,lte=130"`
	}
)

func (m *PersonSimple) Validate() (err error) {
	if strings.Contains(m.Name, "Foo") {
		err = errors.New("name cannot contain 'Foo'")
		return
	}
	return
}

func (m *PersonSimple) ChangeValidationMessage(fieldErr validator.FieldError) (errorMessage string) {
	structFieldName := fieldErr.StructField()
	failedTag := fieldErr.Tag()

	if structFieldName == "Age" && failedTag == "lte" {
		errorMessage = fmt.Sprintf(
			"%s must be less than or equal to %s",
			structFieldName,
			fieldErr.Param(),
		)
	}
	return
}

func NewDefaultPersonSimple() PersonSimple {
	return PersonSimple{
		Name: "John",
		Age:  30,
	}
}
