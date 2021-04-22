package controllers

import (
	"net/http"
	"struttura/config"

	echo "github.com/labstack/echo/v4"
)

func StartWebServer() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(config.Config.Port))
}
