package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var SessionKey = "session"

func SessionMiddleware() echo.MiddlewareFunc {
	return session.Middleware(sessions.NewCookieStore([]byte("secret")))
}

// CheckLoginSession retrun echo.MiddlewareFunc
func CheckLoginSession() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("session", c)
			if err != nil {
				return err
			}

			if isLogin, _ := sess.Values["auth"]; isLogin != true {
				DeleteLoginSession(c)
				return c.Redirect(http.StatusFound, "/login")
			}
			return next(c)
		}
	}
}

// StoreLoginSession return void
func StoreLoginSession(c echo.Context) error {
	sess, _ := session.Get(SessionKey, c)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 0,
		//MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["auth"] = true
	return sess.Save(c.Request(), c.Response())
}

func DeleteLoginSession(c echo.Context) error {
	sess, _ := session.Get(SessionKey, c)
	sess.Options.MaxAge = -1
	return sess.Save(c.Request(), c.Response())
}

//		delete(s.Values, key)
