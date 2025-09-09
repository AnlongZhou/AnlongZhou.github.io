package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Renderer = newTemplate()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
