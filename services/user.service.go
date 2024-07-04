package services

import (
	"github.com/manoelduran/refinancy-with-go/models"
	"github.com/manoelduran/refinancy-with-go/repositories"
)

type UserService struct {

	repository *repositories.UserRepository

}

func NewUserService(repository *repositories.UserRepository) *UserService {

	return &UserService{repository}

}

func (u *UserService) GetUsers() ([]models.User, error) {

	return u.repository.GetUsers()

}

func (u *UserService) GetUser(id uint) (models.User, error) {

	return u.repository.GetUser(id)

}

func (u *UserService) CreateUser(user models.User) (models.User, error) {

	return u.repository.CreateUser(user)

}


func (u *UserService) UpdateUser(id uint, recipe models.User) (models.User, error) {

	return u.repository.UpdateUser(id, recipe)

}

func (r *UserService) DeleteUser(id uint) error {

	return r.repository.DeleteUser(id)

}