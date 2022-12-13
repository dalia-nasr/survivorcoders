package repository

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/matthewhartstonge/argon2"
	"survivorcoders.com/user-go/entity"
	"survivorcoders.com/user-go/utils"
)

func (s UserRepository) CreateToken(user_id uuid.UUID) entity.Token {
	token := entity.Token{UserId: user_id, Type: "invalid", ActivatedAt: time.Now(), ExpiredAt: time.Now().AddDate(1, 0, 0)}
	s.DB.Create(&token)
	s.DB.Save(&token)
	return token
}

func (s UserRepository) CreatePassword(token_id uuid.UUID, upuser *entity.User) (*entity.User, error) {
	token := new(entity.Token)
	if result := s.DB.First(&token, token_id); result.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, result.Error)
	}
	if token.ExpiredAt.Unix() < time.Now().Unix() {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Expired Token")
	}

	user := new(entity.User)
	if result := s.DB.First(&user, token.UserId); result.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, result.Error)
	}
	argon := argon2.DefaultConfig()
	encoded, _ := argon.HashEncoded([]byte(upuser.Password))
	user.Password = string(encoded)
	token.Type = "valid"
	s.DB.Save(user)
	s.DB.Save(token)
	return user, nil
}

func (s UserRepository) Login(userauth *entity.User) (*entity.Token, error) {
	user := new(entity.User)
	if result := s.DB.Where("email = ?", userauth.Email).First(&user); result.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, result.Error)
	}

	authenticated, _ := argon2.VerifyEncoded([]byte(userauth.Password), []byte(user.Password))
	if userauth.Email != user.Email || !authenticated {
		return nil, echo.NewHTTPError(http.StatusNotFound, "wrong email or password")
	}
	jwt, err := utils.CreateJwt(user.Id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	token := new(entity.Token)
	if result := s.DB.Where("user_id = ?", user.Id).First(&token); result.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, result.Error)
	}
	if token.ExpiredAt.Unix() < time.Now().Unix() {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Expired Token")
	}
	token.Jwt = jwt
	s.DB.Save(token)
	return token, nil
}
