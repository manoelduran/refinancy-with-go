package repositories

import (
	"database/sql"

	"github.com/manoelduran/refinancy-with-go/models"
)

type UserRepository struct {
	*GenericRepository[models.User]
}

func NewUserRepository(db *sql.DB) *UserRepository {
	fields := []string{"Name", "Email", "Password"}
	return &UserRepository{
		GenericRepository: NewGenericRepository[models.User](db, "users", fields),
	}
}

func (u *UserRepository) GetUsers() ([]models.User, error) {
	return u.GetAll()
}

func (u *UserRepository) GetUser(id uint) (models.User, error) {
	return u.GetByID(id)
}

func (u *UserRepository) CreateUser(user models.User) (models.User, error) {
	return u.Create(user)
}

func (u *UserRepository) UpdateUser(id uint, user models.User) (models.User, error) {
	return u.Update(id, user)
}

func (u *UserRepository) DeleteUser(id uint) error {
	return u.Delete(id)
}
