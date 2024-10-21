package customvalidator

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	ActualValidator *validator.Validate
}

// New creates the custom validator instance.
func New() *Validator {
	return &Validator{
		ActualValidator: validator.New(),
	}
}

// Validate is the function used to validate any struct.
//
// It has the ability to take in custom validation that
// is made by the developer, specificially using the
// `CustomValidation` interface.
//
// Not limited to only that, it can also replace any
// validation error message with a custom one made by
// you using `CustomValidationMessage` interface.
func (v *Validator) Validate(ptrValue any) (err error) {
	if ptrValue == nil {
		return errors.New("ptrValue cannot be nil")
	}

	typeRef := reflect.TypeOf(ptrValue)
	if typeRef.Kind() != reflect.Ptr {
		return errors.New("ptrValue must be a pointer")
	}
	if typeRef.Elem().Kind() != reflect.Struct {
		return errors.New("ptrValue must be a pointer to a struct")
	}

	// When there's no error we'll use the custom validation
	// made by the developer.
	if err = v.ActualValidator.Struct(ptrValue); err == nil {
		err = v.customValidate(ptrValue)
		return
	}

	// Incase the validation is invalid, we're stopping
	// the process here.
	if _, ok := err.(*validator.InvalidValidationError); !ok {
		return
	}

	validationErrors := err.(validator.ValidationErrors)
	firstFieldError := validationErrors[0]
	err = firstFieldError

	// Try to use custom error message for the validation
	// when the developer decides to overwrite the original error message.
	errWithCustomMessage := v.changeValidationMessage(ptrValue, firstFieldError)
	if errWithCustomMessage != nil {
		err = errWithCustomMessage
	}

	return
}

func (*Validator) customValidate(v any) (err error) {
	customValidation, ok := v.(CustomValidation)
	if !ok {
		return
	}

	err = customValidation.Validate()
	return
}

func (*Validator) changeValidationMessage(v any, fieldErr validator.FieldError) (err error) {
	customValidationMessage, ok := v.(CustomValidationMessage)
	if !ok {
		return
	}

	newErrorMessage := customValidationMessage.ChangeValidationMessage(fieldErr)
	if newErrorMessage == "" {
		return
	}

	err = errors.New(newErrorMessage)
	return
}
