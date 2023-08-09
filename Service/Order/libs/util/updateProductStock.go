package util

import (
	"runtime"
	"sync"
)
var wg sync.WaitGroup
// func UpdateProductStock(productList datatypes.JSON) {
func UpdateProductStock(quantityMapping map[uint]int, token string) (
	error) {
	// var productItems []models.ProductItem
	// // var product_ids_string string

	// byt := []byte(productList)``
	// err := json.Unmarshal(byt, &productItems)
	// if (err != nil) {
	// 	return err
	// }
	
	// for i, v := range productItems {
	// 	// Make a string of id list 
	// 	product_ids_string += fmt.Sprint(product.Id)
	// 	if (i == len(productItems) - 1) {
	// 		break
	// 	}
	// 	product_ids_string += ","
	// }	

	
	// Synchronizing
	// var wg sync.WaitGroup
	productCount := len(quantityMapping)
	runtime.GOMAXPROCS(0)

	// Concurrency begin
	wg.Add(productCount)

	// // Nothing to do with action CREATE
	// var updateItems []models.ProductItem
	// if (action == "DELETE") {

	// }
	// if (action == "UPDATE") {

	// }

	//  ****UPDATE****
	// originalQuantityMapping := make(map[uint]uint)
	// if (action == "UPDATE") {
	// 	// Retrieving order data
	// 	// User Repository instance
	// 	orderRepo := repoimpl.NewOrderRepo(
	// 		driver.MySQL.SQL)
	// 	orders, err := orderRepo.Select("id", orderId)
		
	// 	var productItems []models.ProductItem
	// 	byt := []byte(orders[0].ProductList)
	// 	err := json.Unmarshal(byt, &productItems)
	// 	if (err != nil) {
	// 		return err
	// 	}
	// 	for _, v := range productItems {
	// 		originalQuantityMapping[v.Id] = v.Quantity
	// 	}
	// }

	// for _, v := range productItems {
	// 	go func() {
	// 		// Formatting value to be updated
	// 		// Nothing to do with action CREATE
	// 		var changedValue int
	// 		if (action == "DELETE") {
	// 			changedValue = -(v.Quantity)
	// 		}
	// 		if (action == "UPDATE") {
	// 			// Retrieving order data
	// 			// User Repository instance
	// 			orderRepo := repoimpl.NewOrderRepo(
	// 				driver.MySQL.SQL)
	// 			orders, err := orderRepo.Select("id", paramId)
	// 			changedValue = 
	// 		}

	// 		requestURL := fmt.Sprintf("%s%s%s", 
	// 		baseURL, resource, v.Id)
	// 		bodyString := fmt.Sprintf(`{"change": %d}`, v.Quantity)
	// 		jsonBody := []byte(bodyString)
	// 		bodyReader := bytes.NewReader(jsonBody)

	// 		req, err := http.NewRequest(
	// 			http.MethodPatch, requestURL, bodyReader)
	// 	} ()
	// }
	var err error

	for id, quantity := range quantityMapping {
		// go func() { 
		// 	defer wg.Done()
		// 	// Formatting value to be updated
		// 	// Nothing to do with action CREATE

		// 	// ****DELETE****
		// 	// var changedValue int
		// 	// if (action == "DELETE") {
		// 	// 	changedValue = -(quantity)
		// 	// }
		// 	// if (action == "UPDATE") {
		// 	// 	changedValue = quantity - originalQuantityMapping[id]
		// 	// }
			
		// 	//...
		// 	// Making request to product service
		// 	RequestUpdate(token, id, quantity)
		// 	//...

		// 	fmt.Println(err)
		// } ()

		// Making request to product service
		go RequestUpdate(token, id, quantity)
	}

	// Concurrency end
	wg.Wait()
	
	return err
}