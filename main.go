package main

import (
	"survivorcoders.com/user-go/database"
	"survivorcoders.com/user-go/env"
	"survivorcoders.com/user-go/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	env.EnvVars()

	e := echo.New()

	userController := database.ConnectDb()
	routes.UserRoutes(e, userController)

	e.Logger.Fatal(e.Start(":8000"))
}
