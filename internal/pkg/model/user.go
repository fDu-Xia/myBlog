package model

import (
	"gorm.io/gorm"
	"myBlog/pkg/auth"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
	Nickname string `gorm:"column:nickname"`
	Email    string `gorm:"column:email"`
	Phone    string `gorm:"column:phone"`
}

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Encrypt the user password.
	u.Password, err = auth.Encrypt(u.Password)
	if err != nil {
		return err
	}

	return nil
}
