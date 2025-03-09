package routes

import (
	"demoProject/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// Register the routes
	authenticate := server.Group("/")
	authenticate.Use(middleware.Authenticate)
	authenticate.GET("/events", getEvents)
	authenticate.POST("/events", createEvent)
	authenticate.GET("/events/:id", getEventById)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)
	authenticate.POST("/events/:id/register", registerForEvent)
	authenticate.DELETE("/events/:id/register", unregisterFromEvent)

	// server.GET("/events", getEvents)
	// server.POST("/events", middleware.Authenticate, createEvent)
	// server.GET("/events/:id", getEventById)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
	// Register the routes
}
