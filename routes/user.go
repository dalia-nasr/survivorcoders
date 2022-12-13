package routes

import (
	"github.com/labstack/echo/v4"
	"survivorcoders.com/user-go/controller"
)

func UserRoutes(e *echo.Echo, userController controller.UserController) {

	e.GET("/users", userController.GetAllUsers)
	e.GET("/users/:id", userController.GetUser)
	e.PUT("/users/:id", userController.UpdateUser)
	e.PATCH("/users/:id", userController.UpdatePatchUser)
	e.DELETE("/users/:id", userController.DeleteUser)

	e.POST("/register", userController.RegisterUser)
	e.PATCH("/createpassword/:token", userController.CreatePassword)
	e.GET("/login", userController.Login)
	e.POST("/home", userController.Home)
}
