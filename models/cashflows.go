package models

import (
	"time"
)
type Cashflow struct {
    Id      uint   `json:"id"`
    Title   string `json:"title" validate:"required,min=3,max=15"`
    User_Id  int    `json:"user_id" validate:"required,numeric"`
    Description string `json:"description" validate:"required,min=5,max=20"`
	Earnings int `json:"earnings" validate:"required,numeric"`
	Costs int `json:"costs" validate:"required,numeric"`
    Total int `json:"total" validate:"required,numeric"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}