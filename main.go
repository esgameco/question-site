package main

import (
	"os"

	"github.com/esgameco/question-site/controllers"
	"github.com/esgameco/question-site/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := engine()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func engine() *gin.Engine {
	r := gin.Default()

	// Session handling
	controller := controllers.New()
	r.Use(sessions.Sessions("sessions", cookie.NewStore((models.GetSecret()))))

	// Serve vue.js frontend
	r.Use(static.Serve("/", static.LocalFile("./questionclient/dist", false)))

	// Serve API
	r.POST("/api/login", controller.Login)

	return r
}
