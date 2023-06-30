package customvalidator

import "github.com/go-playground/validator/v10"

// CustomValidation is an interface to allow developers in having
// their own custom validation.
type CustomValidation interface {
	// Validate will be invoked when the actual validation, a.k.a
	// from `validator/v10` doesn't throw any errors.
	//
	// This can be used for having a complex and customized validation
	// based on the developer's preference. Although, not limited to
	// only being validation, it can be used to fill in empty struct
	// fields.
	//
	// Example usage:
	//
	//	// dtos.go
	//	type FindBlogReq struct {
	//		TitleIn  string `json:"titleIn" validate:"required"`
	//		Limit    int64  `json:"limit"`
	//	}
	//
	//	func (d *FindTodosReq) Validate() (err error) {
	//		// Showing custom error
	//		if strings.Contains(d.Title, "Staff") {
	//			err = errors.New("title cannot contain 'Staff' to show staff blogs.")
	//		}
	//		// Changing the value of struct
	//		if d.Limit == 0 || d.Limit > 500 {
	//			d.Limit = 500
	//		}
	//		return
	//	}
	//
	Validate() (err error)
}

// CustomValidationMessage is an interface to allow developers in
// having their own custom error message when it detects something
// wrong in the validation.
type CustomValidationMessage interface {
	// ChangeValidationMessage will be invoked when there's an error
	// in the validation, and then replaces the error message with
	// your own custom error message.
	//
	// It only replaces the error message when it's not empty, as in
	// being empty string like "", otherwise it will use the actual
	// error message from the validation.
	//
	// Example usage:
	//
	//	// dtos.go
	//	type FindBlogReq struct {
	//		TitleIn  string `json:"titleIn" validate:"required"`
	//		Limit    int64  `json:"limit"`
	//	}
	//
	//	func (*FindTodosReq) ChangeValidationMessage(fieldErr validator.FieldError) (errorMessage string) {
	//		jsonFieldName := fieldErr.Field()
	//		failedTag := fieldErr.Tag()
	//
	//		if jsonFieldName == "titleIn" && failedTag == "required" {
	//			errorMessage = fmt.Sprintf("%s must be filled in.", jsonFieldName)
	//		}
	//
	//		return
	//	}
	//
	ChangeValidationMessage(fieldErr validator.FieldError) (errorMessage string)
}
