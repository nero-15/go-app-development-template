package controllers

import (
	"struttura/config"

	echo "github.com/labstack/echo/v4"
)

// StartWebServer is
func StartWebServer() {
	e := echo.New()

	e.GET("/", index)

	e.Logger.Fatal(e.Start(config.Config.Port))
}
