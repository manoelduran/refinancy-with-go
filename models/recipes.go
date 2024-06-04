package models

import (
	"time"
)
type Recipe struct {
    Id    uint   `json:"id" gorm:"primaryKey"`
    Title string `json:"title" validate:"required,min=3,max=15"`
    FromBy string  `json:"from_by" validate:"required,uuid"`
    Description string `json:"description" validate:"required,min=5,max=20"`
    Value int `json:"value" validate:"required,numeric"`
    ReceivedAt time.Time
    ReceivedBy  string    `json:"received_by" validate:"required,uuid"`
    CreatedAt time.Time `json:"created_at" validate:"required,datetime=2006-01-02"`
    UpdatedAt  time.Time `json:"updated_at" validate:"required,datetime=2006-01-02"`
}