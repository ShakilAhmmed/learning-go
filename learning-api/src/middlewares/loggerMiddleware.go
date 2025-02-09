package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		req := c.Request()

		// Log Request Details
		log.Printf("Method: %s | Path: %s | Time: %s", req.Method, req.URL.Path, start.Format(time.RFC3339))

		// Proceed to next handler
		err := next(c)

		// Log Response Time
		log.Printf("Completed in %v", time.Since(start))

		return err
	}
}
