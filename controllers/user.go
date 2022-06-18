package controllers

import (
	"net/http"

	"github.com/esgameco/question-site/models"
	"github.com/gin-gonic/gin"
)

func (con *Controller) Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing email or password"})
		return
	}

	if user.Username == "test" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Logged in",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
	}
}
