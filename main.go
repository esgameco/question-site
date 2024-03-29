package main

import (
	"fmt"
	"os"

	"github.com/esgameco/question-site/controllers"
	"github.com/esgameco/question-site/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("test")
	r := engine()
	r.Use(gin.LoggerWithWriter(os.Stdout))
	r.Use(gin.Recovery())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func engine() *gin.Engine {
	r := gin.New()

	// Session handling
	controller := controllers.New()
	r.Use(sessions.Sessions("sessions", cookie.NewStore(models.GetSecret())))

	// Serve vue.js frontend
	r.Use(static.Serve("/", static.LocalFile("./questionclient/dist", false)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./questionclient/dist/index.html")
	})

	// Serve API
	r.POST("/api/register", controller.Register)
	r.POST("/api/login", controller.Login)
	r.POST("/api/logout", controller.Logout)

	// Authenticated routes
	r.Use(controllers.IsAuthenticated())
	{
		r.GET("/api/questions", controller.GetQuestionsQuery)
		r.GET("/api/questions/all", controller.GetAllQuestions)
		r.GET("/api/question/:id", controller.GetQuestion)
		r.POST("/api/question/create", controller.CreateQuestion)
		r.POST("/api/question/:id/update", controller.UpdateQuestion)
	}

	return r
}
