package main

import (
	"fmt"

	"GoTest/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello World")

	// var x, y int = 1, 2
	// z := 3
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, Echo!")
	// })
	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
