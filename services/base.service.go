package services

import (
	"github.com/manoelduran/refinancy-with-go/repositories"
)
type Service[T any] interface {
    GetAll() ([]T, error)
    GetByID(id uint) (T, error)
    Create(item T) (T, error)
    Update(id uint, item T) (T, error)
    Delete(id uint) error
}


type GenericService[T any] struct {
    repository repositories.Repository[T]
}

func NewGenericService[T any](repository repositories.Repository[T]) *GenericService[T] {
    return &GenericService[T]{repository}
}

func (s *GenericService[T]) GetAll() ([]T, error) {
    return s.repository.GetAll()
}

func (s *GenericService[T]) GetByID(id uint) (T, error) {
    return s.repository.GetByID(id)
}

func (s *GenericService[T]) Create(item T) (T, error) {
    return s.repository.Create(item)
}

func (s *GenericService[T]) Update(id uint, item T) (T, error) {
    return s.repository.Update(id, item)
}

func (s *GenericService[T]) Delete(id uint) error {
    return s.repository.Delete(id)
}