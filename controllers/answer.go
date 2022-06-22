package controllers

import (
	"net/http"

	"github.com/esgameco/question-site/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Get answers
func (con *Controller) GetAnswers(c *gin.Context) {
	var answers []models.Answer
	res := con.Database.Find(&answers)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, answers)
}

// Get answer
func (con *Controller) GetAnswer(c *gin.Context) {
	var answer models.Answer
	res := con.Database.First(&answer, c.Param("id"))

	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answer not found"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// Create answer
func (con *Controller) CreateAnswer(c *gin.Context) {
	var answerInfo models.AnswerInfo

	// Check if json binds
	if err := c.ShouldBindJSON(&answerInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check for empty body
	if answerInfo.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing content"})
		return
	}

	var answer models.Answer

	// Get session
	session := sessions.Default(c)
	userID := session.Get("id")
	answer.AuthorID = int(userID.(uint))

	// Set body
	answer.Body = answerInfo.Body

	// Create answer
	con.Database.Create(&answer)

	c.JSON(http.StatusOK, answer)
}
