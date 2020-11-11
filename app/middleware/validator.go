package middleware

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// CustomValidator return valodate
type CustomValidator struct {
	validator *validator.Validate
}

// Validate return error
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// ValidatorMiddleware return echo.MiddlewareFunc
func ValidatorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Echo().Validator = &CustomValidator{
				validator: validator.New(),
			}
			return next(c)
		}
	}
}
