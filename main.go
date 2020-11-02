package main

import (
	stdLog "log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sonoday8/webapp001/app/config"
	"github.com/sonoday8/webapp001/app/routes"

	"github.com/sonoday8/webapp001/app/env"

	appMiddleware "github.com/sonoday8/webapp001/app/middleware"

	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

// CustomValidator return valodate
type CustomValidator struct {
	validator *validator.Validate
}

// Validate return error
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	loadEnv()
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/session", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["foo"] = "bar"
		sess.Save(c.Request(), c.Response())
		return c.NoContent(http.StatusOK)
	})

	setMiddlewares(e)
	setDebug(e)
	router := routes.Router(e)
	router.Logger.Fatal(router.Start(":" + env.GetStr("SERVER_PORT", "8080")))
}

/**
 * Middleware設定
 */
func setMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(config.AccessLogConfig))
	e.Use(appMiddleware.LoggerMiddleware(env.GetStr("APP_LOG", "logs/application.log")))
	e.Use(middleware.Recover())
	// e.Use(appMiddleware.Sessions())
}

/**
 * .envの読み込み
 */
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		stdLog.Fatal("Error loading .env file")
	}
}

/**
 * デバックモード設定
 */
func setDebug(e *echo.Echo) {
	if env.GetBool("DEBUG", false) {
		e.Debug = true
		e.Logger.SetLevel(log.DEBUG)
	}
}
