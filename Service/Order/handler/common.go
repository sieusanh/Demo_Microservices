package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

// api greeting
func Greeting(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Welcome to Order Service",
	})
}
