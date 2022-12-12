package routes

import (
	"github.com/labstack/echo/v4"
	"survivorcoders.com/user-go/controller"
)

func AuthRoutes(e *echo.Echo, authController controller.AuthController) {
	e.POST("/register", authController.RegisterUser)
	e.PATCH("/createpassword/:token", authController.CreatePassword)
	e.GET("/login", authController.Login)
	e.POST("/home", authController.Home)
}
