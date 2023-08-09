package handler

import (
	libJwt "go-module/libs/jwt"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

func GetAccountInfo(c echo.Context) error {
	
	// Default response 
	resCode := http.StatusOK
	resMes := http.StatusText(resCode)
	
	// Check if account authorized
	request := c.Request()
	authorHeader := request.Header.Get("Authorization")
	splitted := strings.Split(authorHeader, " ")
	if (authorHeader == "" || len(splitted) != 2) {
		resCode = http.StatusForbidden
		resMes := http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}
	
	// Verify access token
	tokenPart := splitted[1]
	token, err := libJwt.Verify(tokenPart)

	if (err != nil || !token.Valid) {
		resCode = http.StatusInternalServerError
		resMes = "Invalid access token."
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}
	
	return c.JSON(resCode, echo.Map{
		"message": resMes,
		"claims": token.Claims,
	})
}
