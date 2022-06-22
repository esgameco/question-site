package models

import "gorm.io/gorm"

type Likes struct {
	gorm.Model
	Amount int `gorm:"type:int;not null;default:0" json:"amount"`
}
