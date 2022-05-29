package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthHandler ...
func AuthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
