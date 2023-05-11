package entity

import (
	"gorm.io/gorm"
)

type User struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Tel			  string    `json:"tel"`
	Zipcode   string    `json:"zipcode"`
	Address   string    `json:"address"`
	Sex				string 	  `json:"sex"`
	gorm.Model
}