package helpers

import (
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/contracts"
	"github.com/labstack/echo/v4"
	"net/http"
)

func JsonResponse(context echo.Context, statusCode int, message string, data interface{}) error {
	response := contracts.ApiResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	return context.JSON(http.StatusOK, response)

}

func JsonErrorResponse(context echo.Context, statusCode int, message string, errors interface{}) error {
	response := contracts.ApiErrorResponse{
		Status:  statusCode,
		Message: message,
		Errors:  errors,
	}
	return context.JSON(statusCode, response)

}
