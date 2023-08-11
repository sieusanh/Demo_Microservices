package util

import (
	"runtime"
	"sync"
)
var wg sync.WaitGroup
func UpdateProductStock(quantityMapping map[uint]int, token string) (
	error) {
	
	// Synchronizing
	productCount := len(quantityMapping)
	runtime.GOMAXPROCS(0)

	// Concurrency begin
	wg.Add(productCount)

	var err error

	for id, quantity := range quantityMapping {
		// Making request to product service
		go RequestUpdate(token, id, quantity)
	}

	// Concurrency end
	wg.Wait()
	
	return err
}