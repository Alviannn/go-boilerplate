package customvalidator

import (
	"errors"

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
	err = v.ActualValidator.Struct(ptrValue)
	// When there's no error we'll use the custom validation
	// made by the developer.
	if err == nil {
		err = v.tryCustomValidation(ptrValue)
		return
	}

	// Incase the validation is invalid, we're stopping
	// the process here.
	_, isOk := err.(*validator.InvalidValidationError)
	if !isOk {
		return
	}

	validationErrors := err.(validator.ValidationErrors)
	firstFieldError := validationErrors[0]

	err = firstFieldError
	// Try to use custom error message for the validation
	// when the developer decides to overwrite the original error message.
	if errWithCustomMessage := v.tryCustomValidationMessage(ptrValue, firstFieldError); errWithCustomMessage != nil {
		err = errWithCustomMessage
	}

	return
}

func (*Validator) tryCustomValidation(v any) (err error) {
	customValidation, isOk := v.(CustomValidation)
	if !isOk {
		return
	}

	err = customValidation.Validate()
	return
}

func (*Validator) tryCustomValidationMessage(v any, fieldErr validator.FieldError) (err error) {
	customValidationMessage, isOk := v.(CustomValidationMessage)
	if !isOk {
		return
	}

	newErrorMessage := customValidationMessage.ChangeValidationMessage(fieldErr)
	if newErrorMessage == "" {
		return
	}

	err = errors.New(newErrorMessage)
	return
}
