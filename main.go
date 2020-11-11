package main

import (
	stdLog "log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sonoday8/webapp001/app/config"
	"github.com/sonoday8/webapp001/app/routes"

	"github.com/sonoday8/webapp001/app/env"

	appcontext "github.com/sonoday8/webapp001/app/context"
	appMiddleware "github.com/sonoday8/webapp001/app/middleware"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

func main() {
	loadEnv()
	e := echo.New()
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
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(appMiddleware.LoggerMiddleware(env.GetStr("APP_LOG", "logs/application.log")))
	e.Use(appMiddleware.ValidatorMiddleware())
	e.Use(appMiddleware.TemplateMiddleware("app/views/*/*.html"))
	// e.Use(appMiddleware.DBConnect())
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(&appcontext.DBContext{c})
		}
	})
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
