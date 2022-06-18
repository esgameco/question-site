package models

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB(connStr string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(connStr), &gorm.Config{})
}

func InitDB(db *gorm.DB) {
	db.AutoMigrate(&Question{}, &Answer{}, &Comment{}, &User{})
}

func GetSecret() []byte {
	return []byte(os.Getenv("SECRET"))
}
