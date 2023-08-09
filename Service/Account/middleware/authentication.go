package middleware

import (
	"github.com/labstack/echo"
	"go-module/libs/custom_type" 
	"go-module/libs/jwt"
	"net/http"
	"strings"

	"fmt"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {		
		// Http error response 
		httpErr := new(echo.HTTPError)

		// Check if account authorized
		request := c.Request()
		authorHeader := request.Header.Get("Authorization")
		splitted := strings.Split(authorHeader, " ")
		if (authorHeader == "" || len(splitted) != 2) {
			httpErr.Code = http.StatusForbidden
			httpErr.Message = "Unauthorized."
			return httpErr
		}

		// Verify access token
		tokenPart := splitted[1]
		token, err := jwt.Verify(tokenPart)
		
		fmt.Println("err: ", err)
		if (err != nil || !token.Valid) {
			httpErr.Code =  http.StatusInternalServerError
			httpErr.Message = "Invalid access token."
			return httpErr
		}

		// Get claims from access token
		claims := token.Claims.(*custom_type.Claims)
		
		// Forwarding claim info to the next middlware
		authenInfo := custom_type.AuthenInfo{}
		authenInfo = claims.AuthenInfo
		c.Set("Role", authenInfo.Role)

		// Go to next 
		err = next(c)
		return err
	}
}
