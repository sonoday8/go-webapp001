package middleware

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

// Sessions return echo.MiddlewareFunc
func Sessions() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//セッションを設定
			store := sessions.NewCookieStore([]byte("secret-key"))
			//セッション保持時間
			store.MaxAge(86400)
			err := next(c)
			// after
			return err
		}
	}
}
