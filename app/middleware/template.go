package middleware

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template return strucr
type Template struct {
	templates *template.Template
}

// Render return error
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// TemplateMiddleware return echo.MiddlewareFunc
func TemplateMiddleware(templatePath string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			t := &Template{
				templates: template.Must(template.ParseGlob(templatePath)),
			}
			c.Echo().Renderer = t
			// after
			return next(c)
		}
	}
}
