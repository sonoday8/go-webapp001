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

	appContext "github.com/sonoday8/webapp001/app/context"
	appHandler "github.com/sonoday8/webapp001/app/handler"
	appMiddleware "github.com/sonoday8/webapp001/app/middleware"
)

func main() {
	loadEnv()
	e := echo.New()
	setMiddlewares(e)
	setDebug(e)
	e.HTTPErrorHandler = appHandler.CustomHTTPErrorHandler
	router := routes.Router(e)
	router.Logger.Fatal(router.Start(":" + env.GetStr("SERVER_PORT", "8080")))
}

/**
 * Middleware設定
 */
func setMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(config.AccessLogConfig))
	e.Use(middleware.Recover())
	e.Use(appMiddleware.SessionMiddleware())
	e.Use(appMiddleware.LoggerMiddleware(env.GetStr("APP_LOG", "logs/application.log")))
	e.Use(appMiddleware.ValidatorMiddleware())
	e.Use(appMiddleware.TemplateMiddleware("app/views/*.html", "app/views/*/*.html"))
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(&appContext.DBContext{c})
		}
	})
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:token",
	}))
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
