package entities

import "time"

type (
 InsertUserDto struct {
  Id        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
  Amount    uint32    `json:"amount"`
  CreatedAt time.Time `json:"createdAt"`
 }

 User struct {
  Id        uint32    `json:"id"`
  Amount    uint32    `json:"amount"`
  CreatedAt time.Time `json:"createdAt"`
 }

 UserPushNotificationDto struct {
  Title        string `json:"title"`
  Amount       uint32 `json:"amount"`
  ReportedTime string `json:"createdAt"`
 }
)