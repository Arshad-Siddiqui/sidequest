package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint    `json:"id" gorm:"primaryKey"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Quests   []Quest `json:"quests" gorm:"foreignKey:User_id"`
}
