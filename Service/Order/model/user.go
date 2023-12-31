package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	AccountId primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email string			     `json:"email" bson:"email"`
	Password string				 `json:"password" bson:"password"`
	Name string			 		 `json:"displayName" bson:"displayName"`
} 		
