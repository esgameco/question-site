package controllers

import (
	"net/http"

	"github.com/esgameco/question-site/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (con *Controller) GetQuestions(c *gin.Context) {
	var questions []models.Question
	res := con.Database.Find(&questions)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (con *Controller) GetQuestion(c *gin.Context) {
	var question models.Question
	res := con.Database.Preload("Author").First(&question, c.Param("id"))

	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	c.JSON(http.StatusOK, question)
}

func (con *Controller) CreateQuestion(c *gin.Context) {
	var question models.Question

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	userID := session.Get("id")
	question.AuthorID = int(userID.(uint))

	// Create question
	con.Database.Create(&question)

	// Add question to user
	var user models.User
	if err := con.Database.First(&user, userID); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		con.Database.Delete(&question)
		return
	}
	con.Database.Model(&user).Association("Questions").Append(&question)

	c.JSON(http.StatusOK, question)
}

func (con *Controller) UpdateQuestion(c *gin.Context) {
	var question models.Question

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	con.Database.Save(&question)

	c.JSON(http.StatusOK, question)
}
