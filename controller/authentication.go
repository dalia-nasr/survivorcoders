package controller

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"survivorcoders.com/user-go/entity"
	"survivorcoders.com/user-go/repository"
	"survivorcoders.com/user-go/utils"
)

type AuthController struct {
	AuthRepository repository.AuthRepository
	UserRepository repository.UserRepository
}

func (r AuthController) RegisterUser(c echo.Context) error {
	User := &entity.User{}
	if err := c.Bind(User); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	user := r.UserRepository.CreateUser(User)
	token := r.AuthRepository.CreateToken(user.Id)
	email := utils.SendMail(token.Id, user.Email)

	if email != nil {
		return c.JSON(http.StatusNotFound, email)
	}
	return c.JSON(http.StatusCreated, user)
}

func (r AuthController) CreatePassword(c echo.Context) error {
	token := c.Param("token")
	casttoken, _ := uuid.FromString(token)
	User := &entity.User{}
	if err := c.Bind(User); err != nil {
		return err
	}
	result, err := r.AuthRepository.CreatePassword(casttoken, User)
	if err == nil {
		return c.JSON(http.StatusOK, result)
	}
	return c.JSON(http.StatusNotFound, err)
}

func (r AuthController) Login(c echo.Context) error {
	User := &entity.User{}
	if err := c.Bind(User); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	token, err := r.AuthRepository.Login(User)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, token)
}

func (r AuthController) Home(c echo.Context) error {
	tokenheader := c.Request().Header.Get("Token")
	token, err := utils.VerifyToken(tokenheader)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userid, ok := claims["user_id"].(string)
		if !ok {
			return err
		}
		id, _ := uuid.FromString(userid)
		if result, err := r.UserRepository.GetUser(id); err == nil {
			return echo.NewHTTPError(http.StatusOK, "Hello "+result.Name)
		} else {
			return c.JSON(http.StatusNotFound, err.Error())
		}
	}
	return err
}
