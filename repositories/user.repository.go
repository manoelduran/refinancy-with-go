package repositories

import (
	"github.com/manoelduran/refinancy-with-go/models"
	"gorm.io/gorm"
)
type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {

	return &UserRepository{db}

}

func (u *UserRepository) GetUsers() ([]models.User, error) {

	var users []models.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil

}

func (u *UserRepository) GetUser(id uint) (models.User, error) {

	var user models.User
	result := u.db.First(&user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil

}

func (u *UserRepository) CreateUser(user models.User) (models.User, error) {

	result := u.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil

}

func (u *UserRepository) UpdateUser(id uint, user models.User) (models.User, error) {

	result := u.db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil

}
func (u *UserRepository) DeleteUser(id uint) error {

	result := u.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil

}