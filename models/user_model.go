package models

/*
Create a User struct with required properties.
We added omitempty and validate:"required" to the struct tag to tell Gin to ignore
empty fields and make the field required, respectively.
*/
type User struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
	Title    string `json:"title,omitempty" validate:"required"`
}
