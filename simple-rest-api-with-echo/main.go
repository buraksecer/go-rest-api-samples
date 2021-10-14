package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/test", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

//localhost:8080/test
// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}