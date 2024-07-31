package models

import (
	"time"
)
type Recipe struct {
    Id      uint   `json:"id"`
    Title   string `json:"title" validate:"required,min=3,max=15"`
    From_By  int    `json:"from_by" validate:"required,numeric"`
    Description string `json:"description" validate:"required,min=5,max=20"`
    Value int `json:"value" validate:"required,numeric"`
    Received_By  int    `json:"received_by" validate:"required,numeric"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}