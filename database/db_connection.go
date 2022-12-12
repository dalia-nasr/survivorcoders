package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"survivorcoders.com/user-go/controller"
	"survivorcoders.com/user-go/repository"
)

func ConnectDb() (controller.UserController, controller.AuthController) {
	dbConnection, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_LINK")), &gorm.Config{})
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	userController := controller.UserController{
		UserRepository: repository.UserRepository{
			DB: dbConnection,
		},
	}
	authController := controller.AuthController{
		AuthRepository: repository.AuthRepository{
			DB: dbConnection,
		},
	}
	return userController, authController
}
