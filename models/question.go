package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255);not null" json:"title"`
	Body     string `gorm:"type:text;not null" json:"body"`
	Likes    int    `gorm:"type:int;not null;default:0" json:"likes"`
	Answers  []Answer
	Comments []Comment
	Author   User
}
