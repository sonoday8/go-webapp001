package middleware

import (
	"github.com/labstack/echo/v4"
	appContext "github.com/sonoday8/webapp001/app/context"
)

func DBConnect() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(&appContext.DBContext{c})
		}
	}
}
