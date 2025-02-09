package models

import (
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (user *User) Validate() map[string]string {
	return helpers.FormatValidation(user)
}
