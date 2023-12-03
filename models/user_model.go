package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
Create a User struct with required properties.
We added omitempty and validate:"required" to the struct tag to tell Gin to ignore empty fields and make the field required, respectively.
*/
type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}
