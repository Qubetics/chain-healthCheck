package routes

import (
	"healthCheck/handlers"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	e.GET("/health", handlers.HealthCheck)
}
