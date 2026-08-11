package models

import "time"

type Todo struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar(255);not null" json:"title"`
	Done bool `gorm:"default:false" json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
