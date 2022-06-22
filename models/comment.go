package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body     string     `gorm:"type:text;not null" json:"body"`
	Likes    Likes      `gorm:"foreignkey:LikesID" json:"likes"`
	LikesID  int        `json:"-"`
	Replies  []*Comment `gorm:"many2many:comment_replies" json:"replies"`
	Author   *User      `gorm:"foreignkey:AuthorID" json:"author"`
	AuthorID int        `json:"author_id"`
}
