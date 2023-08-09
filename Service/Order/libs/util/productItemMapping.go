package util

import (
	"gorm.io/datatypes"
	"encoding/json"
	"fmt"
	models "go-module/model"
)

func ProductItemMapping(productList datatypes.JSON) (
	map[uint]uint, string, error) {
	
	var productItems []models.ProductItem
	var product_ids_string string
	quantityMapping := make(map[uint]uint)

	byt := []byte(productList)
	err := json.Unmarshal(byt, &productItems)
	if (err != nil) {
		return nil, "", err
	}

	for i, v := range productItems {
		// Quantity mapping
		quantityMapping[v.Id] = v.Quantity

		// Make id list string
		product_ids_string += fmt.Sprint(v.Id)
		if (i == len(productItems) - 1) {
			break
		}
		product_ids_string += ","
	}
	return quantityMapping, product_ids_string, err
}
