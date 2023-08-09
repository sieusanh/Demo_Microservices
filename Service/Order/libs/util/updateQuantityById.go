package util

import (
	"go-module/config/product_service"
	"net/http"
	"bytes"
	"fmt"
)

// Retrieving product data
func RequestUpdate(token string, id uint, quantity int) {
	defer wg.Done()

	baseURL := product_service.HOST
	resource := product_service.ROUTE_UPDATE_STOCK
	requestURL := fmt.Sprintf("%s%s%d", 
	baseURL, resource, id)
	bodyString := fmt.Sprintf(`{"change": %d}`, quantity)
	jsonBody := []byte(bodyString)
	bodyReader := bytes.NewBuffer(jsonBody)

	// create new HTTP PATCH request with JSON payload
	req, err := http.NewRequest(
		http.MethodPatch, requestURL, bodyReader)

	// set content-type header to JSON
	authenHeader := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authenHeader)

	// create HTTP client and execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}