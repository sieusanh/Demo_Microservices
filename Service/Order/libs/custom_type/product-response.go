package custom_type

import "go-module/model"

type ProductResponse struct {
	Total uint				`json:"total"`
	Data []model.Product	`json:"data"`
}
	