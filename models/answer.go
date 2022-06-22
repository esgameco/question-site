package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Body       string     `gorm:"type:text;not null" json:"body"`
	Likes      Likes      `gorm:"foreignkey:LikesID" json:"likes"`
	LikesID    int        `json:"-"`
	Comments   []*Comment `gorm:"many2many:answer_comments" json:"comments"`
	Author     *User      `gorm:"foreignkey:AuthorID" json:"author"`
	AuthorID   int        `json:"-"`
	Question   *Question  `gorm:"foreignkey:QuestionID" json:"question"`
	QuestionID int        `json:"-"`
}

type AnswerInfo struct {
	Body       string `json:"body"`
	QuestionID int    `json:"question_id"`
}
