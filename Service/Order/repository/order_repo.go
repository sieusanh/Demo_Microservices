package repository

import (
	models "go-module/model"
)

type OrderRepo interface {
	Select(field, value string) ([]models.Order, error)
	Insert(o models.Order) (error)
	UpdateById(id string, 
		payload models.Order) error
	DeleteById(id string) error 
	CreateTable(string) error
	DropTable(string) error
}
