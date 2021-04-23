package controllers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "struttura",
		"age":  10,
	})
}
