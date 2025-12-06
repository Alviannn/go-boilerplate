package hooks

import "go-boilerplate/pkg/customerror"

func ErrorMarshallerHook(err error) any {
	customError, ok := err.(*customerror.Error)
	if !ok {
		return err
	}
	return customError.ToJSON()
}

