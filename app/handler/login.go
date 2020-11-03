package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type viewParams struct {
	LoginForm *LoginForm
	Name      string
}

// LoginIndex return echo.HandlerFunc
func LoginIndex(c echo.Context) error {
	loginForm := new(LoginForm)
	data := viewParams{
		LoginForm: loginForm,
		Name:      "sonoda",
	}
	return c.Render(http.StatusOK, "login_index.html", data)
}

// LoginForm return form
type LoginForm struct {
	LoginID  string `form:"loginID" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

// LoginExec return
func LoginExec(c echo.Context) error {
	loginForm := new(LoginForm)
	var err = c.Bind(loginForm)
	if err != nil {
		return err
	}
	err = c.Validate(loginForm)
	if err != nil {
		c.Logger().Debug("Valid error")
		loginForm.Password = "" //Password clear
		data := viewParams{
			LoginForm: loginForm,
			Name:      "sonoda",
		}
		//		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		return c.Render(http.StatusUnprocessableEntity, "login_index.html", data)
	}

	data := viewParams{
		LoginForm: loginForm,
		Name:      "sonoda",
	}

	return c.Render(http.StatusOK, "login_index.html", data)
}
