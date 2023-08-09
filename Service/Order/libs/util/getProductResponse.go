package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	models "go-module/model"
	"go-module/libs/custom_type"
	"go-module/config/product_service"

	"fmt"
)

func GetProductResponse(product_ids_string string) (
	[]models.Product, error) {
		
	// Retrieving product data
	baseURL := product_service.HOST
	resource := product_service.ROUTE_GET
	requestURL := fmt.Sprintf("%s%s%s", 
		baseURL, resource, product_ids_string)
	resp, err := http.Get(requestURL)
	if (err != nil) {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if (err != nil) {
		return nil, err
	}

	product_response := custom_type.ProductResponse{}
	byt := []byte(string(body))
	err = json.Unmarshal(byt, &product_response)

	return product_response.Data, err
}
