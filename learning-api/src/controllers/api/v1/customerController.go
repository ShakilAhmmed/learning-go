package v1

import (
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/contracts"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomerController struct {
}

func (customerController *CustomerController) Index(context echo.Context) error {
	customers := []map[string]interface{}{
		{"id": 101, "name": "John Doe", "email": "john@example.com", "age": 30},
		{"id": 102, "name": "Alice Smith", "email": "alice@example.com", "age": 25},
		{"id": 103, "name": "Bob Brown", "email": "bob@example.com", "age": 40},
	}

	response := contracts.ApiResponse{
		Status:  http.StatusOK,
		Message: "Customers fetched successfully.",
		Data:    customers,
	}

	return context.JSON(http.StatusOK, response)
}
