package repositories

import (
	"gorm.io/gorm"
)

type Repository[T any] interface {
    GetAll() ([]T, error)
    GetByID(id uint) (T, error)
    Create(item T) (T, error)
    Update(id uint, item T) (T, error)
    Delete(id uint) error
}

type GenericRepository[T any] struct {
    db *gorm.DB
}

func NewGenericRepository[T any](db *gorm.DB) *GenericRepository[T] {
    return &GenericRepository[T]{db}
}

func (r *GenericRepository[T]) GetAll() ([]T, error) {
    var items []T
    result := r.db.Find(&items)
    if result.Error != nil {
        return nil, result.Error
    }
    return items, nil
}

func (r *GenericRepository[T]) GetByID(id uint) (T, error) {
    var item T
    result := r.db.First(&item, id)
    if result.Error != nil {
        return item, result.Error
    }
    return item, nil
}

func (r *GenericRepository[T]) Create(item T) (T, error) {
    result := r.db.Create(&item)
    if result.Error != nil {
        return item, result.Error
    }
    return item, nil
}

func (r *GenericRepository[T]) Update(id uint, item T) (T, error) {
    result := r.db.Model(&item).Where("id = ?", id).Updates(item)
    if result.Error != nil {
        return item, result.Error
    }
    return item, nil
}

func (r *GenericRepository[T]) Delete(id uint) error {
    var item T
    result := r.db.Delete(&item, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}