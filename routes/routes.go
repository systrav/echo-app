package routes

import (
	h "GoTest/internal/handlers"
	s "GoTest/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	g := e.Group("/api")
	// userService := services.NewUserService()
	// userController := handlers.NewUserController(userService)
	// g.GET("/hello", userController.Hello)
	g.GET("/hello", h.NewUserController(s.NewUserService()).Hello)
}
