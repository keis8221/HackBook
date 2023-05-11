package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Users 	 []User    `gorm:"many2many:user_categories;"`
	gorm.Model
}