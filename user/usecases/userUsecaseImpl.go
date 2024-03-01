package usecases

import (
	"time"

	"tansan/user/entities"
	"tansan/user/models"
	"tansan/user/repositories"
)

type userUsecaseImpl struct {
 userRepository repositories.UserRepository
 userMessaging  repositories.UserMessaging
}

func NewUserUsecaseImpl(
 userRepository repositories.UserRepository,
 userMessaging repositories.UserMessaging,
) UserUsecase {
 return &userUsecaseImpl{
  userRepository: userRepository,
  userMessaging:  userMessaging,
 }
}

func (u *userUsecaseImpl) UserDataProcessing(in *models.AddUserData) error {
 insertUserData := &entities.InsertUserDto{
  Amount: in.Amount,
 }

 if err := u.userRepository.InsertUserData(insertUserData); err != nil {
  return err
 }

 pushUserData := &entities.UserPushNotificationDto{
  Title:        "User Detected 🪳 !!!",
  Amount:       in.Amount,
  ReportedTime: time.Now().Local().Format("2006-01-02 15:04:05"),
 }

 if err := u.userMessaging.PushNotification(pushUserData); err != nil {
  return err
 }

 return nil
}