package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// CheckLoginSession retrun echo.MiddlewareFunc
func CheckLoginSession() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("session", c)
			if err != nil {
				return err
			}

			if isLogin, _ := sess.Values["auth"]; isLogin != true {
				return c.Redirect(http.StatusFound, "/login")
			}
			return next(c)
		}
	}
}

// StoreLoginSession return void
func StoreLoginSession(c echo.Context) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["auth"] = true
	sess.Save(c.Request(), c.Response())
}
