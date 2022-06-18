package controllers

import (
	"log"

	"github.com/esgameco/question-site/models"
	"gorm.io/gorm"
)

type Controller struct {
	Database *gorm.DB
}

func New() *Controller {
	db, err := models.GetDB("question-site.db")
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}
	models.InitDB(db)

	return &Controller{db}
}
