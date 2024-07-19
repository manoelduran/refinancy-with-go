package services

import (
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/repositories"
)

type UserService struct {
	*GenericService[models.User]

}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
        GenericService: NewGenericService(repository),
    }
}

func (u *UserService) GetUsers() ([]models.User, error) {

	return u.repository.GetAll()

}

func (u *UserService) GetUser(id uint) (models.User, error) {

	return u.repository.GetByID(id)

}

func (u *UserService) CreateUser(user models.User) (models.User, error) {

	return u.repository.Create(user)

}


func (u *UserService) UpdateUser(id uint, recipe models.User) (models.User, error) {

	return u.repository.Update(id, recipe)

}

func (r *UserService) DeleteUser(id uint) error {

	return r.repository.Delete(id)

}