package middleware

import (
	"github.com/labstack/echo"
	"net/http"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {		
		role := c.Get("Role")

		if (role != "Admin") {
			httpErr := new(echo.HTTPError)
			httpErr.Code = http.StatusForbidden
			httpErr.Message = "Unauthorized."
			
			return httpErr
		}

		err := next(c)
		return err
	}
}
