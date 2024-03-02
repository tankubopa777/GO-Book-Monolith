package repositories

import "tansan/user/entities"

type UserRepository interface {
 InsertUserData(in *entities.InsertUserDto) error
 FindAll() ([]*entities.User, error)
}