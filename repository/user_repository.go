package repository

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"survivorcoders.com/user-go/entity"
)

type UserRepository struct {
	DB *gorm.DB
}

func (s UserRepository) GetAll() []entity.User {
	var users []entity.User
	s.DB.Find(&users)
	return users
}

func (s UserRepository) GetUser(userid uuid.UUID) (*entity.User, error) {
	user := new(entity.User)
	if result := s.DB.First(&user, userid); result.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "email, password, and name required")
	}
	return user, nil
}

func (s UserRepository) CreateUser(user *entity.User) *entity.User {
	s.DB.Create(&user)
	s.DB.Save(&user)
	return user
}

func (s UserRepository) UpdateUser(id uuid.UUID, upuser *entity.User) (*entity.User, error) {
	user := new(entity.User)
	if result := s.DB.First(&user, id); result.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, result.Error)
	}
	user.Email = upuser.Email
	user.Name = upuser.Name
	s.DB.Save(user)
	return user, nil
}

func (s UserRepository) UpdatePatchUser(id uuid.UUID, upuser *entity.User) (*entity.User, error) {
	user := new(entity.User)
	if result := s.DB.First(&user, id); result.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, result.Error)
	}
	if upuser.Email != "" {
		user.Email = upuser.Email
	}
	if upuser.Name != "" {
		user.Name = upuser.Name
	}
	s.DB.Save(user)
	return user, nil
}

func (s UserRepository) DeleteUser(id uuid.UUID) error {
	db := s.DB.Delete(&entity.User{}, id)
	if db.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Ensure id is correct")
	}
	return nil
}
