package repository

import (
	models "go-module/model"
	"go-module/libs/custom_type"
)

type UserRepo interface {
	Count() (uint, error)
	FindAll() ([]custom_type.ConciseData, error)
	FindByField(field, value string) (models.User, error) 
	UpdateById(id string, payload models.User) error
	CheckLoginInfo(email string, password string) (models.User, error)
	Insert(u models.User) error 
	RemoveAll() error
}