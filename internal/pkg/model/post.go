package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Username string `gorm:"column:username;not null"`
	PostID   string `gorm:"column:post_id;not null"`
	Title    string `gorm:"column:title;not null"`
	Content  string `gorm:"column:content"`
}
