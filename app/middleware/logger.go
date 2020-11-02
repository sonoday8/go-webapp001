package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
)

// LoggerMiddleware return echo.MiddlewareFunc
func LoggerMiddleware(fileName string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Echo().HideBanner = true
			c.Echo().HidePort = true
			logFile, logErr := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			if logErr != nil {
				panic("log file error")
			}
			c.Echo().Logger.SetOutput(logFile)

			err := next(c)
			// after
			return err
		}
	}
}
