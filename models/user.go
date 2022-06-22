package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string      `gorm:"type:varchar(255);not null;unique" json:"username"`
	Email     string      `gorm:"type:varchar(255);not null;unique" json:"-"`
	Password  string      `gorm:"type:varchar(255);not null" json:"-"`
	Questions []*Question `gorm:"many2many:user_questions" json:"questions"`
}

type UserInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
