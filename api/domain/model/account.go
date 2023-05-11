package entity

import (
	"gorm.io/gorm"
)

type Account struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	UserId		uint      `json:"user_id"`
	gorm.Model
}