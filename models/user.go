package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null;unique" json:"username" binding:"required"`
	Email    string `gorm:"type:varchar(255);not null;unique" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
}
