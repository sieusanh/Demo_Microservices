package model

type ProductItem struct {
	Id uint		`json:"id"`
	Quantity uint	`json:"quantity"`
}

type UpdateItem struct {
	Id int		`json:"id"`
	Quantity uint	`json:"quantity"`
}
