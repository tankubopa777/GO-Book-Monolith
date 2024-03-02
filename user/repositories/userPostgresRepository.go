package repositories

import (
	"tansan/user/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
 db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) UserRepository {
 return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) InsertUserData(in *entities.InsertUserDto) error {
 data := &entities.User{
  Amount: in.Amount,
 }

 result := r.db.Create(data)

 if result.Error != nil {
  log.Errorf("InsertUserData: %v", result.Error)
  return result.Error
 }

 log.Debugf("InsertUserData: %v", result.RowsAffected)
 return nil
}

func (r *userPostgresRepository) FindAll() ([]*entities.User, error) {
    var users []*entities.User
    result := r.db.Find(&users)
    if result.Error != nil {
        log.Errorf("FindAll: %v", result.Error)
        return nil, result.Error
    }
    return users, nil
}