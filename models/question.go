package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title    string     `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	Body     string     `gorm:"type:text;not null" json:"body" binding:"required"`
	Likes    Likes      `gorm:"foreignkey:LikesID" json:"likes"`
	LikesID  int        `json:"-"`
	Answers  []*Answer  `gorm:"many2many:question_answers" json:"answers"`
	Comments []*Comment `gorm:"many2many:question_comments" json:"comments"`
	Author   *User      `gorm:"foreignkey:AuthorID" json:"author"`
	AuthorID int        `json:"-"`
}

type QuestionsQuery struct {
	SortType   string `json:"sortType"`
	Page       int    `json:"page"`
	NumPerPage int    `json:"numPerPage"`
}
