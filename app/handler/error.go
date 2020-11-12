package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type errorViewParams struct {
	Error string
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	data := errorViewParams{
		Error: err.Error(),
	}
	if err := c.Render(code, fmt.Sprintf("error%d.html", code), data); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
