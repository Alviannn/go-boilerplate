package customvalidator

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	ActualValidator *validator.Validate
	helper          helper
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
func (v *Validator) Validate(value any) (err error) {
	if value == nil {
		return errors.New("value cannot be nil")
	}

	if !v.helper.isStruct(value) {
		err = errors.New("value must be a struct or a pointer to a struct")
		return
	}

	// When there's no error we'll use the custom validation
	// made by the developer.
	if err = v.ActualValidator.Struct(value); err != nil {
		// Incase the validation is invalid, we're stopping
		// the process here.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return
		}

		validationErrors := err.(validator.ValidationErrors)
		firstFieldError := validationErrors[0]
		err = firstFieldError

		// Try to use custom error message for the validation
		// when the developer decides to overwrite the original error message.
		errWithCustomMessage := v.changeValidationMessage(value, firstFieldError)
		if errWithCustomMessage != nil {
			err = errWithCustomMessage
		}
		if err != nil {
			return
		}
		return
	}

	if err = v.customValidate(value); err != nil {
		return
	}
	if err = v.recursiveValidation(value); err != nil {
		return
	}
	return
}

func (v *Validator) recursiveValidation(value any) (err error) {
	elemValueRef := reflect.Indirect(reflect.ValueOf(value))

	if v.helper.isSliceOrArray(value) {
		for i := 0; i < elemValueRef.Len(); i++ {
			err = v.handleRecursiveValidationForFieldOrIndex(elemValueRef.Index(i))
			if err != nil {
				return
			}
		}
		return
	}

	if v.helper.isStruct(value) {
		for i := 0; i < elemValueRef.NumField(); i++ {
			err = v.handleRecursiveValidationForFieldOrIndex(elemValueRef.Field(i))
			if err != nil {
				return
			}
		}
	}

	return
}

func (v *Validator) handleRecursiveValidationForFieldOrIndex(valueRef reflect.Value) (err error) {
	var valueToPass any
	if valueRef.Kind() == reflect.Ptr || v.helper.isSliceOrArray(valueRef.Interface()) {
		valueToPass = valueRef.Interface()
	} else {
		valueToPass = valueRef.Addr().Interface()
	}

	if v.helper.isSliceOrArray(valueToPass) {
		err = v.recursiveValidation(valueToPass)
		return
	}
	if v.helper.isStruct(valueToPass) {
		err = v.Validate(valueToPass)
		return
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
