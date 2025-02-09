package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterGlobalMiddlewares(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(LoggerMiddleware, TestMiddleware)
}
