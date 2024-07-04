package models

import (
	"time"
)
type User struct {
    Id    uint   `json:"id" gorm:"primaryKey"`
    Name string `json:"name" validate:"required,min=5,max=20"`
    Email string  `json:"email"`
    Password string `json:"password" validate:"required,min=4,max=10"`
    CreatedAt time.Time `json:"created_at" `
    UpdatedAt  time.Time `json:"updated_at" `
}