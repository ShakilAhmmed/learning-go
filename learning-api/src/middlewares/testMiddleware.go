package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
)

func TestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("Hello from custom middleware")
		return next(c)
	}
}
