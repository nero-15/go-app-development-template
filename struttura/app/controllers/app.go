package controllers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
