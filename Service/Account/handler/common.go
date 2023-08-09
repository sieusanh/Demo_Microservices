package handler

import (
	"encoding/json"
	mongoConfig "go-module/config/mongodb"
	driver "go-module/driver"
	models "go-module/model"
	libJwt "go-module/libs/jwt"
	"go-module/repository/repoimpl"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Greeting(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Welcome to User Service",
	})
}

func Register(c echo.Context) error {
	
	// Request instance
	request := c.Request()

	// Default response 
	resCode := http.StatusOK
	resMes := http.StatusText(resCode)

	// Decode payload json into payload
	var payload models.User
	err := json.NewDecoder(request.Body).Decode(&payload)

	// Error response
	if (err != nil) {
		resCode = http.StatusBadRequest
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	// User Repository instance
	userRepo := repoimpl.NewUserRepo(
		driver.Mongo.Client.
		Database(mongoConfig.DB_NAME))

	// Check if this account existed
	_, err = userRepo.FindByField("email", payload.Email)
	if err == nil {
		resCode = http.StatusConflict
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	} 

	// Get current user number count
	user_count, err := userRepo.Count()

	// Insert new account
	payload.Id = strconv.FormatUint(uint64(user_count + 1), 10)
	err = userRepo.Insert(payload)
	
	if err != nil {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	// Generate new access token
	tokenString := ""
	tokenString, err = libJwt.GenToken(payload) 

	// Check if generate token failed
	if err != nil {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	return c.JSON(resCode, echo.Map{
		"message": resMes,
		"token": tokenString,
	})
}

func Login(c echo.Context) error {

	// Request instance
	request := c.Request()

	// Default response 
	resCode := http.StatusOK
	resMes := http.StatusText(resCode)
	
	// Decode payload json into payload
	var payload models.LoginData
	err := json.NewDecoder(request.Body).Decode(&payload)
	
	// Error response
	if (err != nil) {
		resCode = http.StatusBadRequest
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	// User Repository instance
	userRepo := repoimpl.NewUserRepo(
		driver.Mongo.Client.
		Database(mongoConfig.DB_NAME))
	
	// Check account info
	var user models.User
	user, err = userRepo.CheckLoginInfo(
		payload.Email, 
		payload.Password,
	)

	if err != nil {
		resCode = http.StatusUnauthorized
		resMes = "Incorrect username or password."
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}
	
	// Generate new access token
	tokenString, err := libJwt.GenToken(user) 

	// Error response
	if err != nil {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	return c.JSON(resCode, echo.Map{
		"message": resMes,
		"token": tokenString,
	})	
}
