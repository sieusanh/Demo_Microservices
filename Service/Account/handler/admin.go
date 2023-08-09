package handler

import (
	"github.com/labstack/echo"
	"encoding/json"
	mongoConfig "go-module/config/mongodb"
	"go-module/repository/repoimpl"
	"go-module/libs/custom_type" 
	models "go-module/model"
	"go-module/libs/time"
	"go-module/driver"
	"strconv"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"

	"fmt"
)

func count() (uint, error) {
	// User Repository instance
	userRepo := repoimpl.NewUserRepo(
		driver.Mongo.Client.
		Database(mongoConfig.DB_NAME))

	// Get current user number count
	user_count, err := userRepo.Count()
	return user_count, err
}

func Count(c echo.Context) error {

	// Default response 
	resCode := http.StatusOK
	resMes := http.StatusText(resCode)

	// Making database request
	user_count, err := count()

	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	return c.JSON(resCode, echo.Map{
		"data": user_count,
	})
}

func Find(c echo.Context) error {

	// Default response 
	resCode := http.StatusOK
	resMes := http.StatusText(resCode)

	// User Repository instance
	userRepo := repoimpl.NewUserRepo(
		driver.Mongo.Client.
		Database(mongoConfig.DB_NAME))

	res, err := userRepo.FindAll()
	
	if (err != nil) {
		resCode = http.StatusNotFound
		resMes = http.StatusText(resCode)
		return c.JSON(404, echo.Map{
			"message": resMes,
		})
	}

	return c.JSON(resCode, echo.Map{
		"data": res,
	})
}

func FindById(c echo.Context) error {

	field := "id"
	value := c.Param("id")

	// Default response 
	resCode := http.StatusOK
	resMes := http.StatusText(resCode)

	// User Repository instance
	userRepo := repoimpl.NewUserRepo(
		driver.Mongo.Client.
		Database(mongoConfig.DB_NAME))
		
	res, err := userRepo.FindByField(field, value)
	
	if (err != nil) {
		resCode = http.StatusNotFound
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	return c.JSON(resCode, echo.Map{
		"data": res,
	})
}

func UpdateById(c echo.Context) error {

	// Request params
	// val := c.QueryParam("value")	
	request := c.Request()
	body := request.Body
	id := c.Param("id")

	// Default response
	resCode := http.StatusNoContent
	resMes := ""

	var payload models.User
	err := json.NewDecoder(body).Decode(&payload)
	
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
		
	err = userRepo.UpdateById(id, payload)

	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	return c.JSON(resCode, echo.Map{})
}

func RemoveAll(c echo.Context) error {
	// Default response 
	resCode := http.StatusOK
	resMes := http.StatusText(resCode)

	// User Repository instance
	userRepo := repoimpl.NewUserRepo(
		driver.Mongo.Client.
		Database(mongoConfig.DB_NAME))

	err := userRepo.RemoveAll()
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	return c.JSON(resCode, echo.Map{
		"message": resMes,
	})
}

// Populating data
func PopulatingData(c echo.Context) error {
	// https://dummyjson.com/users	

	// Default response 
	resCode := http.StatusOK
	resMes := http.StatusText(resCode)

	// Making database request
	user_count, err := count()

	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	count_threshold := uint(10)
	if (user_count > count_threshold) {
		resCode = http.StatusNotAcceptable
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	// Time tracking
	defer time.Timer()

	// Retrieving product data
	getUrl := "https://dummyjson.com/users?limit=100"                
	resp, err := http.Get(getUrl)                                                     
	body, err := ioutil.ReadAll(resp.Body)

	// // Format response data
	// var user_response custom_type.UserResponse
	var user_response custom_type.DummyResponse
	byt := []byte(string(body))
	err = json.Unmarshal(byt, &user_response)
	fmt.Println(err)

	// User Repository instance
	userRepo := repoimpl.NewUserRepo(
		driver.Mongo.Client.
		Database(mongoConfig.DB_NAME))
	
	// Synchronizing
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(0)

	// Concurrency begin
	wg.Add(100)

	for _, v := range user_response.Users {
		insert_data := models.User{}
		insert_data.Id = strconv.FormatUint(uint64(v.Id), 10)
		insert_data.Name = v.Name
		insert_data.Age = v.Age
		insert_data.Gender = v.Gender
		insert_data.Phone = v.Phone
		insert_data.Email = v.Email
		insert_data.Username = v.Username
		insert_data.Password = v.Password
		insert_data.Role = v.Role

		// err = userRepo.Insert(insert_data)
		// With concurrency
		go func() {
			defer wg.Done()
			err = userRepo.Insert(insert_data)
			fmt.Println(insert_data)
		} ()
	}

	// Concurrency end
	wg.Wait()

	fmt.Println("err: ", err)
	return c.JSON(resCode, echo.Map{
		"message": resMes,
	})
}