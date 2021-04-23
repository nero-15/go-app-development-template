package controllers

import (
	"html/template"
	"io"
	"struttura/config"

	echo "github.com/labstack/echo/v4"
)

// Template ...
type Template struct {
	templates *template.Template
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// StartWebServer is used echo
func StartWebServer() {
	e := echo.New()
	renderer := &Template{
		templates: template.Must(template.ParseGlob("app/views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", index)

	e.Logger.Fatal(e.Start(config.Config.Port))
}
