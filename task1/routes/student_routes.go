package routes

import (
	"github.com/labstack/echo/v4"
	"survivorcoders.com/task1/controller"
)

func Routes(e *echo.Echo, studentController controller.StudentController) {

	e.GET("/students", studentController.GetAllStudents)
	e.GET("/students/:id", studentController.GetStudent)
	e.POST("/students", studentController.CreateStudent)
	e.PUT("/students/:id", studentController.UpdateStudent)
	e.PATCH("/students/:id", studentController.UpdatePatchStudent)
	e.DELETE("/students/:id", studentController.DeleteStudent)
}
