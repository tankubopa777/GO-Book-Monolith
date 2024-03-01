package repositories

import "tansan/user/entities"

type UserMessaging interface {
 PushNotification(m *entities.UserPushNotificationDto) error
}

