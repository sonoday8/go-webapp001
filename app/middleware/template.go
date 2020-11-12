package middleware

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"

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
func TemplateMiddleware(templatePatterns ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var templatePaths []string
			for _, templatePattern := range templatePatterns {
				filenames, err := filepath.Glob(templatePattern)
				if err != nil {
					return err
				}
				if len(filenames) == 0 {
					return fmt.Errorf("html/template: pattern matches no files: %#q", templatePattern)
				}
				templatePaths = append(templatePaths, filenames...)
			}
			temp, err := template.ParseFiles(templatePaths...)
			if err != nil {
				return err
			}
			t := &Template{
				//templates: template.Must(template.ParseGlob(templatePath)),
				templates: temp,
			}
			c.Echo().Renderer = t
			// after
			return next(c)
		}
	}
}
