package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body    string `gorm:"type:text;not null" json:"body"`
	Likes   int    `gorm:"type:int;not null;default:0" json:"likes"`
	Replies []Comment
	Author  User
}
