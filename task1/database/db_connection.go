package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"survivorcoders.com/task1/controller"
	"survivorcoders.com/task1/repository"
)

func ConnectDb() controller.StudentController {
	dbConnection, err := gorm.Open(postgres.Open("postgres://postgres:admin@localhost:5432/students"), &gorm.Config{})
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	studentController := controller.StudentController{
		StudentRepository: repository.StudentRepository{
			DB: dbConnection,
		},
	}
	return studentController
}
