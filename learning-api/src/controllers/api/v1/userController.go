package v1

import (
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/contracts"
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/controllers"
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/models"
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

	response := userController.Response(http.StatusOK, "Users fetched successfully", users)

	return context.JSON(http.StatusOK, response)
}

func (userController *UserController) Store(context echo.Context) error {

	var user models.User

	// Bind JSON request body to the user struct
	if err := context.Bind(&user); err != nil {
		return context.JSON(http.StatusBadRequest, contracts.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request format",
			Data:    nil,
		})
	}

	if validationErrors := user.Validate(); validationErrors != nil {
		return context.JSON(http.StatusBadRequest, contracts.ApiErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  validationErrors, // Returns validation errors
		})
	}

	return context.JSON(http.StatusOK, user)
}
