package handler

import (
	"github.com/labstack/echo/v4"
	appModels "github.com/sonoday8/webapp001/app/models"
	"net/http"
)

// LoginForm return form
type SignupForm struct {
	LoginID  string `form:"loginID" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type signupViewParams struct {
	Form  SignupForm
	Error string
	Text  string
}

// いらないかも
func SignupIndex(c echo.Context) error {
	form := new(SignupForm)

	var user = appModels.User{LoginID: form.LoginID}
	exists, err := appModels.ExistsUser(c, user)
	if err != nil {
		return err
	}
	c.Logger().Debug(exists)
	return c.Render(http.StatusOK, "signup_index.html", nil)
}

func SignupExec(c echo.Context) error {
	form := new(SignupForm)
	if err := c.Bind(form); err != nil {
		return err
	}
	var user = appModels.User{LoginID: form.LoginID, Password: form.Password}
	exists, err := appModels.ExistsUser(c, user)
	if err != nil {
		return err
	}
	if exists {
		return c.Render(http.StatusOK, "signup_index.html",
			&signupViewParams{
				Form:  *form,
				Error: "メールアドレスは既に登録済み",
			})
	}

	done, err := appModels.CreateUser(c, user)
	if err != nil {
		c.Logger().Error(err)
	}
	if done == false {
		return c.Render(http.StatusOK, "signup_index.html", signupViewParams{
			Form:  *form,
			Error: "登録できず",
		})
	}
	return c.Render(http.StatusOK, "signup_comp.html", nil)
}
