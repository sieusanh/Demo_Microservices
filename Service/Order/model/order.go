package model
import "gorm.io/datatypes"

type Order struct {
	Id int 						`json:"id"`
	TotalPrice float32			`json:"totalPrice"`
	UserId string				`json:"userId"`
	// ProductList: map[string]int	`json:"productList"`
	// ProductList []ProductItem	`json:"productList"`
	ProductList datatypes.JSON  `json:"productList"`
}
