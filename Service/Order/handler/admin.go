package handler

import (
	"github.com/labstack/echo"
	"gorm.io/datatypes"
	"encoding/json"
	"net/http"
	"fmt"
	"go-module/repository/repoimpl"
	"go-module/libs/util"
	models "go-module/model"
	"go-module/driver"
)

func GetRespById(paramId string) ([]models.Order, error) {
		
	// Retrieving order data
	// User Repository instance
	orderRepo := repoimpl.NewOrderRepo(
		driver.MySQL.SQL)
	// Being used for both api get all and get by id
	orders, err := orderRepo.Select("id", paramId)
	if (err != nil) {
		return nil, err
	}

	var resp_orders []models.Order

	// Loop through every order item
	for _, order := range orders {                                  
	
		quantityMapping, product_ids_string, err := 
			util.ProductItemMapping(order.ProductList)
		products, err := 
			util.GetProductResponse(product_ids_string)
		fmt.Println(err)
		for i, v := range products {
			products[i].Quantity = quantityMapping[v.Id]
		} 

		jsonProducts, err := json.Marshal(&products)
		order.ProductList = datatypes.JSON(jsonProducts)
		resp_orders = append(resp_orders, order)
		fmt.Println(err)
	}

	return resp_orders, err
}

// Read

func Find(c echo.Context) error {   
	// Request param
	paramId := c.Param("id")

	// Default response 
	resCode := http.StatusOK
	resMes := ""

	// Get data
	resp_orders, err := GetRespById(paramId)
	
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	return c.JSON(resCode, echo.Map{
		"total": len(resp_orders),
		"data": resp_orders,
	})
}

// Delete
func DeleteById(c echo.Context) error {

	// Request params
	paramId := c.Param("id")
	tokenString := c.Get("Token")

	// Default response
	resCode := http.StatusNoContent
	resMes := ""

	// Get mapping of product item
	resp_orders, err := GetRespById(paramId)
	quantityMapping, _, err := 
		util.ProductItemMapping(resp_orders[0].ProductList)

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
		quantityIntMaping[id] = int(quantity)
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

	// DELETE DATA
	// Order Repository instance
	orderRepo := repoimpl.NewOrderRepo(
		driver.MySQL.SQL)
	err = orderRepo.DeleteById(paramId)

	// Error response
	if (err != nil) {
		resCode = http.StatusInternalServerError
		resMes = http.StatusText(resCode)
		return c.JSON(resCode, echo.Map{
			"message": resMes,
		})
	}

	return c.JSON(200, echo.Map{
		"code": 200,
		"message": "ok",
	})
}