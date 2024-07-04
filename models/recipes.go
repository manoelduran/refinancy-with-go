package models

import (
	"time"
)
type Recipe struct {
    Id      uint   `json:"id" gorm:"primaryKey"`
    Title   string `json:"title" validate:"required,min=3,max=15"`
    FromBy  int    `json:"from_by" validate:"required,numeric"`
    Description string `json:"description" validate:"required,min=5,max=20"`
    Value int `json:"value" validate:"required,numeric"`
    ReceivedAt  time.Time `json:"received_at"`
    ReceivedBy  int    `json:"received_by" validate:"required,numeric"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}