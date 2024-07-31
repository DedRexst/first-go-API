package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", postEvents)
	authenticated.PUT("/events/:id", putEvents)
	authenticated.DELETE("/events/:id", deleteEvents)
	authenticated.POST("/events/:id/register", postRegister)
	authenticated.DELETE("/events/:id/register", deleteRegister)

	server.POST("/signup", postUser)
	server.POST("/login", login)
}
