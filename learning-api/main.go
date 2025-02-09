package main

import (
	"fmt"
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/middlewares"
	"github.com/ShakilAhmmed/learning-go/llearning-api/src/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello From Api")
	e := echo.New()

	middlewares.RegisterGlobalMiddlewares(e)
	routes.RegisterRoutes(e)
	e.Start(":8080")
}
