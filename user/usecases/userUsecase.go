package usecases

import "tansan/user/models"

type UserUsecase interface {
 UserDataProcessing(in *models.AddUserData) error
}