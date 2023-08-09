package handler

import (
	"github.com/labstack/echo"
	"encoding/json"
	"net/http"
	"go-module/repository/repoimpl"
	models "go-module/model"
	"go-module/driver"
	"go-module/libs/util"
)

// Create
func Create(c echo.Context) error {

	// Request instance
	request := c.Request()
	tokenString := c.Get("Token")

	// Default response 
	resCode := http.StatusOK
	resMes := http.StatusText(resCode)

	// Decoding param data
	payload := models.Order{}
	err := json.NewDecoder(request.Body).Decode(&payload)

	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}
	
	// Get product item mapping data
	quantityMapping, product_ids_string, err := 
		util.ProductItemMapping(payload.ProductList)
	
	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	// UPDATE PRODUCT STOCK
	// Copy a map with different value type
	quantityIntMaping := make(map[uint]int)
	for id, quantity := range quantityMapping {
		quantityIntMaping[id] = -int(quantity)
	}

	err = util.UpdateProductStock(
		quantityIntMaping, tokenString.(string))

	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	// CALCULATING TOTAL PRICE
	var calculatedPrice float32 = 0
	products, err := 
		util.GetProductResponse(product_ids_string)
	for _, v := range products {
		discountPrice := v.Price - (v.Price * v.Discount / 100)
		quantity := float32(quantityMapping[v.Id]) 
		calculatedPrice += discountPrice * quantity
	}
	payload.TotalPrice = calculatedPrice
	
	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}
                              
	// INSERT DATA
	// Order Repository instance
	// orderRepo := repoimpl.NewOrderRepo(db.SQL)
	orderRepo := repoimpl.NewOrderRepo(
		driver.MySQL.SQL)
	err = orderRepo.Insert(payload)

	// Error response
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

// Update
func UpdateById(c echo.Context) error {

	// Request params
	request := c.Request()
	body := request.Body
	paramId := c.Param("id")
	tokenString := c.Get("Token")

	// Default response
	resCode := http.StatusNoContent
	resMes := ""

	// Decoding param data
	var payload models.Order
	err := json.NewDecoder(body).Decode(&payload)
	
	// Error response
	if (err != nil) {
		resCode = http.StatusBadRequest
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	// Get mapping of product item
	resp_orders, err := GetRespById(paramId)
	quantityMapping, product_ids_string, err := 
		util.ProductItemMapping(resp_orders[0].ProductList)
	
	// Get mapping of update product item
	updateQuantityMapping, _, err := 
		util.ProductItemMapping(payload.ProductList)

	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}	

	// UPDATE PRODUCT STOCK
	// Copy a map with different value type
	quantityIntMaping := make(map[uint]int)
	for id, quantity := range quantityMapping {
		quantityIntMaping[id] = int(quantity) - int(updateQuantityMapping[id])
	}

	err = util.UpdateProductStock(quantityIntMaping, tokenString.(string))

	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	// CALCULATING TOTAL PRICE
	var calculatedPrice float32 = 0
	products, err := 
		util.GetProductResponse(product_ids_string)
	for _, v := range products {
		discountPrice := v.Price - (v.Price * v.Discount / 100)
		quantity := float32(updateQuantityMapping[v.Id]) 
		calculatedPrice += discountPrice * quantity
	}
	payload.TotalPrice = calculatedPrice
	
	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	// UPDATE DATA
	orderRepo := repoimpl.NewOrderRepo(
		driver.MySQL.SQL)
	err = orderRepo.UpdateById(paramId, payload)

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
