package controllers

import (
	"net/http"

	"github.com/esgameco/question-site/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Checks if user is logged in
func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("id")

		if user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		c.Next()
	}
}

// Registers user
func (con *Controller) Register(c *gin.Context) {
	var userInfo models.UserInfo

	// Checks if login details are valid
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check for empty email, username or password
	if userInfo.Email == "" || userInfo.Username == "" || userInfo.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing email, username or password"})
		return
	}

	// Check if username is taken
	var user models.User
	user.Username = userInfo.Username
	user.Email = userInfo.Email
	user.Password = userInfo.Password

	res := con.Database.First(&user, "username = ? OR email = ?", user.Username, user.Email)
	if res.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email already taken"})
		return
	}

	// Register user
	con.Database.Create(&user)

	// Creates session
	session := sessions.Default(c)
	session.Set("id", user.ID)
	session.Save()

	// Logged in
	c.JSON(http.StatusOK, gin.H{
		"message": "Registered",
	})
}

// Creates session for user
func (con *Controller) Login(c *gin.Context) {
	var user models.User
	var userInfo models.UserInfo

	// Checks if login details are valid
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check for empty username or password
	if userInfo.Username == "" || userInfo.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing username or password"})
		return
	}

	// Get user from database
	res := con.Database.First(&user, "username = ?", userInfo.Username)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
		return
	}

	// Check if password is correct
	if userInfo.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	// Creates session
	session := sessions.Default(c)
	session.Set("id", user.ID)
	session.Save()

	// Logged in
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in",
	})
}

func (con *Controller) Logout(c *gin.Context) {
	// Removes session
	session := sessions.Default(c)
	user := session.Get("id")

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
		return
	}

	session.Delete("id")
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out",
	})
}
