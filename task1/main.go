package main

import (
	"survivorcoders.com/task1/database"
	"survivorcoders.com/task1/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	studentController := database.ConnectDb()
	routes.Routes(e, studentController)

	e.Logger.Fatal(e.Start(":5000"))
}
