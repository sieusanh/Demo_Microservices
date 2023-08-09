package model

// import (
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

type User struct {
	// Id primitive.ObjectID	`json:"id,omitempty" bson:"id,omitempty"`
	Id string				`json:"id,omitempty" bson:"id,omitempty"`
	Name string				`json:"name,omitempty" bson:"name,omitempty"`
	Age uint				`json:"age,omitempty" bson:"age,omitempty"`
	Gender string			`json:"gender,omitempty" bson:"gender,omitempty"`
	Phone string			`json:"phone,omitempty" bson:"phone,omitempty"`
	Email string			`json:"email,omitempty" bson:"email,omitempty"`
	Username string			`json:"username,omitempty" bson:"username,omitempty"`
	Password string			`json:"password,omitempty" bson:"password,omitempty"`
	Role string				`json:"role,omitempty" bson:"role,omitempty"`
} 		

type DummyData struct {
	Id uint					`json:"id,omitempty" bson:"id,omitempty"`
	Name string				`json:"firstName,omitempty" bson:"name,omitempty"`
	Age uint				`json:"age,omitempty" bson:"age,omitempty"`
	Gender string			`json:"gender,omitempty" bson:"gender,omitempty"`
	Phone string			`json:"phone,omitempty" bson:"phone,omitempty"`
	Email string			`json:"email" bson:"email"`
	Username string			`json:"username" bson:"username"`
	Password string			`json:"password" bson:"password"`
	Role string				`json:"role,omitempty" bson:"role,omitempty"`
}
