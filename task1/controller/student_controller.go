package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"survivorcoders.com/task1/entity"
	"survivorcoders.com/task1/repository"
)

type StudentController struct {
	StudentRepository repository.StudentRepository
}

func (r StudentController) GetAllStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, r.StudentRepository.GetAll())
}

func (r StudentController) GetStudent(c echo.Context) error {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	} else if result, err := r.StudentRepository.GetStudent(id); err == nil {
		return c.JSON(http.StatusOK, result)
	} else {
		return c.JSON(http.StatusNotFound, err.Error())
	}
}

func (r StudentController) CreateStudent(c echo.Context) error {
	student := &entity.Student{}
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusCreated, r.StudentRepository.CreateStudent(student))
}

func (r StudentController) DeleteStudent(c echo.Context) error {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(http.StatusOK, r.StudentRepository.DeleteStudent(id))
	}
}

func (r StudentController) UpdateStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	student := &entity.Student{}
	if err := c.Bind(student); err != nil {
		return err
	}
	result, _ := r.StudentRepository.UpdateStudent(id, student)
	return c.JSON(http.StatusOK, result)
}

func (r StudentController) UpdatePatchStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	student := &entity.Student{}
	if err := c.Bind(student); err != nil {
		return err
	}
	result, _ := r.StudentRepository.UpdatePatchStudent(id, student)
	return c.JSON(http.StatusOK, result)
}
