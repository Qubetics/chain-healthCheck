package app

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func StartServer(e *echo.Echo) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}
	log.Printf("Health check service running on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
