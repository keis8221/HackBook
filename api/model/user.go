package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Tel			  string    `json:"tel"`
	Zipcode   string    `json:"zipcode"`
	Address   string    `json:"address"`
	Sex				string 	  `json:"sex"`
	Categories []Category `gorm:"many2many:user_categories;"`
}