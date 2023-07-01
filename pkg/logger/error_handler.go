package logger

import "go-boilerplate/pkg/responses"

type CustomErrorInLog struct {
	Error       string   `json:"error"`
	SourceError string   `json:"sourceError"`
	Stack       []string `json:"stack"`
}

func HandleCustomError(err error) any {
	customError, ok := err.(*responses.CustomError)
	if !ok {
		return err.Error()
	}

	return CustomErrorInLog{
		Error:       customError.Error(),
		SourceError: customError.GetWorkingError().Error(),
		Stack:       customError.GetStackTrace(),
	}
}
