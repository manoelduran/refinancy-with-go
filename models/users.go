package models

import (
	"time"
)
type User struct {
    Id    uint   `json:"id" gorm:"primaryKey"`
    Name string `json:"name" validate:"required,min=5,max=20"`
    Email string  `json:"email" validate:"required, email"`
    Password string `json:"description" validate:"required,min=4,max=10"`
    Value int `json:"value" validate:"required,numeric"`
    CreatedAt time.Time `json:"created_at" validate:"required,datetime=2006-01-02"`
    UpdatedAt  time.Time `json:"updated_at" validate:"required,datetime=2006-01-02"`
}