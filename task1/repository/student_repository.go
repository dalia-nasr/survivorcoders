package repository

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"survivorcoders.com/task1/entity"
)

type StudentRepository struct {
	DB *gorm.DB
}

func (s StudentRepository) GetAll() []entity.Student {
	var students []entity.Student
	s.DB.Find(&students)
	return students
}

func (s StudentRepository) GetStudent(id int) (*entity.Student, error) {
	student := new(entity.Student)
	if result := s.DB.First(&student, id); result.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "first and last name required")
	}
	return student, nil
}

func (s StudentRepository) CreateStudent(student *entity.Student) *entity.Student {
	s.DB.Create(&student)
	s.DB.Save(&student)
	return student
}

func (s StudentRepository) DeleteStudent(id int) error {
	db := s.DB.Delete(&entity.Student{}, id)
	if db.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Ensure id is correct")
	}
	return nil
}

func (s StudentRepository) UpdateStudent(id int, upstudent *entity.Student) (*entity.Student, error) {
	student := new(entity.Student)
	if result := s.DB.First(&student, id); result.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, result.Error)
	}
	student.FirstName = upstudent.FirstName
	student.LastName = upstudent.LastName
	s.DB.Save(student)
	return student, nil
}

func (s StudentRepository) UpdatePatchStudent(id int, upstudent *entity.Student) (*entity.Student, error) {
	student := new(entity.Student)
	if result := s.DB.First(&student, id); result.Error != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, result.Error)
	}
	if upstudent.FirstName != "" {
		student.FirstName = upstudent.FirstName
	}
	if upstudent.LastName != "" {
		student.LastName = upstudent.LastName
	}
	s.DB.Save(student)
	return student, nil
}
