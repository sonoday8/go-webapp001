package handler

import (
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	appContext "github.com/sonoday8/webapp001/app/context"
	appMiddleware "github.com/sonoday8/webapp001/app/middleware"
	appModels "github.com/sonoday8/webapp001/app/models"
)

// Login Page params
type viewParams struct {
	LoginForm *LoginForm
	Token     string
	Title     string
}

// LoginIndex return echo.HandlerFunc
func LoginIndex(c echo.Context) error {
	loginForm := new(LoginForm)
	k, err := c.Cookie(middleware.DefaultCSRFConfig.CookieName)
	if err != nil {
		return err
	}
	data := viewParams{
		LoginForm: loginForm,
		Token:     k.Value,
		Title:     "ログイン",
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
	if err := c.Bind(loginForm); err != nil {
		return err
	}

	if err := c.Validate(loginForm); err != nil {
		loginForm.Password = "" //Password clear
		data := viewParams{
			LoginForm: loginForm,
		}
		return c.Render(http.StatusUnprocessableEntity, "login_index.html", data)
	}

	password := []byte(loginForm.Password)
	user := new(appModels.User)
	user.LoginID = loginForm.LoginID
	dbc := c.(*appContext.DBContext)

	if err := dbc.DBConn(func(db *gorm.DB) error {
		db.Select("password").Find(&user)
		return nil
	}); err != nil {
		c.Logger().Error(err)
		return err
	}

	var passwdHash []byte
	passwdHash = []byte(user.Password)

	if err := bcrypt.CompareHashAndPassword(passwdHash, password); err != nil {
		c.Logger().Error(err)
		return err
	}
	if err := appMiddleware.StoreLoginSession(c); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/")
}
