package controller

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"

	"survivorcoders.com/user-go/entity"
	"survivorcoders.com/user-go/repository"
)

type UserController struct {
	UserRepository repository.UserRepository
}

func (r UserController) GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, r.UserRepository.GetAll())
}

func (r UserController) GetUser(c echo.Context) error {
	if userid, err := uuid.FromString(c.Param("id")); err != nil {
		return c.JSON(http.StatusNotFound, "missing params")
	} else if result, err := r.UserRepository.GetUser(userid); err == nil {
		return c.JSON(http.StatusOK, result)
	} else {
		return c.JSON(http.StatusNotFound, err.Error())
	}
}

func (r UserController) DeleteUser(c echo.Context) error {
	if userid, err := uuid.FromString(c.Param("id")); err != nil {
		return c.JSON(http.StatusNotFound, "missing params")
	} else {
		return c.JSON(http.StatusOK, r.UserRepository.DeleteUser(userid))
	}
}

func (r UserController) UpdateUser(c echo.Context) error {
	userid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "missing params")
	}
	User := &entity.User{}
	if err := c.Bind(User); err != nil {
		return err
	}
	result, _ := r.UserRepository.UpdateUser(userid, User)
	return c.JSON(http.StatusOK, result)
}

func (r UserController) UpdatePatchUser(c echo.Context) error {
	userid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "missing params")
	}
	User := &entity.User{}
	if err := c.Bind(User); err != nil {
		return err
	}
	result, _ := r.UserRepository.UpdatePatchUser(userid, User)
	return c.JSON(http.StatusOK, result)
}
