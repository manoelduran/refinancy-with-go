package models

type Recipe struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Value string `json:"value"`
}