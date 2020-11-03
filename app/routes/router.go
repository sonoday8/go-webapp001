package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sonoday8/webapp001/app/handler"
)

// Router return echo.Echo
func Router(e *echo.Echo) *echo.Echo {
	e.GET("/login", handler.LoginIndex)
	e.POST("/login", handler.LoginExec)
	e.File("/signup", "public/signup.html")

	root := e.Group("/")
	// root.Use(middleware.KeyAuthWithConfig(config.AuthConfig))
	// root.Use(middleware.JWTWithConfig(config.AuthConfig))
	root.File("", "public/index.html")

	// e.POST("/signup", handler.Signup)
	// e.File("/login", "public/login.html")
	// e.POST("/login", handler.Login)
	// e.File("/todos", "public/todos.html")

	// api := e.Group("/api")
	// api.Use(middleware.JWTWithConfig(handler.Config))
	// api.GET("/todos", handler.GetTodos)
	// api.POST("/todos", handler.AddTodo)
	// api.DELETE("/todos/:id", handler.DeleteTodo)
	// api.PUT("/todos/:id/completed", handler.UpdateTodo)

	return e
}
