package requests

import (
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/helpers"
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/models"
)

type UserRequest struct {
	models.User
	Password string `json:"password" validate:"required"`
}

func (userRequest *UserRequest) Validate() map[string]string {
	return helpers.FormatValidation(userRequest)
}
