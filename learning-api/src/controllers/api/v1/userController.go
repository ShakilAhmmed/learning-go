package v1

import (
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/controllers"
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/helpers"
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/models"
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/requests"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	controllers.Controller
}

func (userController *UserController) Index(context echo.Context) error {

	users := []models.User{
		{
			Id:       1,
			Name:     "John Doe",
			Email:    "john@example.com",
			Age:      30,
			Password: "john@example.com",
		},
		{
			Id:       2,
			Name:     "Alice Smith",
			Email:    "alice@example.com",
			Age:      25,
			Password: "alice@example.com",
		},
		{
			Id:       3,
			Name:     "Bob Brown",
			Email:    "bob@example.com",
			Age:      40,
			Password: "bob@example.com",
		},
	}

	return helpers.JsonResponse(context, http.StatusOK, "Users fetched successfully", users)
}

func (userController *UserController) Store(context echo.Context) error {

	//var userRequest requests.UserRequest

	userRequest := new(requests.UserRequest)
	// Bind JSON request body to the user request
	if err := context.Bind(userRequest); err != nil {
		return helpers.JsonResponse(context, http.StatusBadRequest, "Invalid request format", nil)
	}

	if validationErrors := userRequest.Validate(); validationErrors != nil {
		return helpers.JsonErrorResponse(context, http.StatusBadRequest, "Validation failed", validationErrors)
	}

	return helpers.JsonResponse(context, http.StatusOK, "User created successfully", nil)
}
