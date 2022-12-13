package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"survivorcoders.com/user-go/controller"
	"survivorcoders.com/user-go/repository"
)

func ConnectDb() controller.UserController {
	dblink := os.Getenv("POSTGRES_LINK")
	dbConnection, err := gorm.Open(postgres.Open(dblink), &gorm.Config{})
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	userController := controller.UserController{
		UserRepository: repository.UserRepository{
			DB: dbConnection,
		},
	}
	return userController
}
