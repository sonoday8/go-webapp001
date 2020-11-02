package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sonoday8/webapp001/app/services"
)

// Index return error
func Index(c echo.Context) error {
	c.Logger().Debug("AAAAAAAA")
	return c.File("public/login.html")
}

// Login return error
func Login(c echo.Context) error {
	c.Logger().Debug("BB")
	services.Login(c)
	return c.File("public/login.html")
}
