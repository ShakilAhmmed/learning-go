package routes

import (
	v1 "github.com/ShakilAhmmed/learning-go/llearning-api/src/controllers/api/v1"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	api := e.Group("/api/v1")

	users := api.Group("/users")
	userController := new(v1.UserController)
	users.GET("", userController.Index)
	users.POST("", userController.Store)

	customers := api.Group("/customers")
	customerController := new(v1.CustomerController)
	customers.GET("", customerController.Index)

}
