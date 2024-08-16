package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
	Nickname string `gorm:"column:nickname"`
	Email    string `gorm:"column:email"`
	Phone    string `gorm:"column:phone"`
}
