package logger

import "go-boilerplate/pkg/responses"

func HandleCustomError(err error) any {
	customError, ok := err.(*responses.CustomError)
	if !ok {
		return err
	}

	return customError.BuildResponse()
}
