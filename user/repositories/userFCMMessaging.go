package repositories

import (
	"tansan/user/entities"

	"github.com/labstack/gommon/log"
)
   
   type userFCMMessaging struct{}
   
   func NewUserFCMMessaging() UserMessaging {
	return &userFCMMessaging{}
   }
   
   func (c *userFCMMessaging) PushNotification(m *entities.UserPushNotificationDto) error {
	// ... handle logic to push FCM notification here ...
	log.Debugf("Pushed FCM notification with data: %v", m)
   
	return nil
   }