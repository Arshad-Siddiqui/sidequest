package models

import (
	"gorm.io/gorm"
)

type Quest struct {
	gorm.Model
	Id          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User_id     uint   `json:"user_id" gorm:"foreignKey:Id"`
}
