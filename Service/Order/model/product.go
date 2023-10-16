package model

type Product struct {
	// Id string			`json:"id"`
	Id uint				`json:"id"`
	Title string		`json:"title,omitempty"`
	Price float32		`json:"price,omitempty"`
    Discount float32	`json:"discount,omitempty"`
	Quantity uint		`json:"quantity,omitempty"`
}
