package main

import (
	"survivorcoders.com/user-go/database"
	"survivorcoders.com/user-go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	userController, authController := database.ConnectDb()
	routes.UserRoutes(e, userController)
	routes.AuthRoutes(e, authController)

	e.Logger.Fatal(e.Start(":8000"))
}
