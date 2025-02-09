package controllers

import (
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/contracts"
)

type Controller struct {
}

func (controller *Controller) Response(statusCode int, message string, data interface{}) contracts.ApiResponse {
	return contracts.ApiResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}

}
