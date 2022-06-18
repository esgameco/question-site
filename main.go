package main

import (
	"github.com/esgameco/question-site/handlers"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve vue.js frontend
	r.Use(static.Serve("/", static.LocalFile("./questionclient/dist", false)))

	// Serve API
	r.POST("/api/login", handlers.Login)

	r.Run(":8080")
}
