package main

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve vue.js frontend
	r.Use(static.Serve("/", static.LocalFile("./questionclient/dist", false)))

	// Serve API
	r.GET("/api/endpoint", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test",
		})
	})

	r.Run(":8080")
}
