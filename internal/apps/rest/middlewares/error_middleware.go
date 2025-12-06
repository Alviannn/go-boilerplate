package middlewares

import (
	"context"
	"errors"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/response"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func CustomErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed { // Prevent sending double response
			return
		}

		customErr := customerror.New().
			WithContext(c.Request().Context()).
			WithSourceError(err)

		if currentCustomErr, ok := err.(*customerror.Error); ok {
			isOverrideError := currentCustomErr.IsPanic || errors.Is(currentCustomErr.GetWorkingError(), context.DeadlineExceeded)
			if isOverrideError {
				customErr = currentCustomErr
			}
		}

		log.Error().
			Ctx(c.Request().Context()).
			Err(customErr).
			Msg(customErr.Message)

		res := response.NewBuilder().
			WithError(customErr).
			Build()

		c.JSON(res.StatusCode, res.Data)
	}
}
