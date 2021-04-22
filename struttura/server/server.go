package server

import (
	"struttura/app/controllers"
	"struttura/config"

	echo "github.com/labstack/echo/v4"
)

// Start is
func Start() {
	e := echo.New()

	e.GET("/", controllers.Index)

	e.Logger.Fatal(e.Start(config.Config.Port))
}
