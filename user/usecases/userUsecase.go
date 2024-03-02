package usecases

import (
	"tansan/user/entities" // Add this import statement
	"tansan/user/models"
)

type UserUsecase interface {
    UserDataProcessing(in *models.AddUserData) error
    GetAllUsers() ([]*entities.User, error) // Now entities.User will be recognized
}