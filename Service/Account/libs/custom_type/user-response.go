package custom_type

import "go-module/model"

type UserResponse struct {
	Total uint				`json:"total"`
	Users []model.User		`json:"users"`
}
	
type DummyResponse struct {
	Total uint				`json:"total"`
	Users []model.DummyData	`json:"users"`
}
	
type ConciseData struct {
	Id string		`json:"id,omitempty" bson:"id,omitempty"`
	Name string		`json:"name,omitempty" bson:"name,omitempty"`
	Username string	`json:"username" bson:"username"`
}