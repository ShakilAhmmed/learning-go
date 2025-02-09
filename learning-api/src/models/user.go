package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"required"`
	Password string `json:"-" validate:"required"`
}
