package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Get all events.
	server.GET("/events", getEvents)

	// Get an event by ID.
	server.GET("/event/:id", getEvent)

	// Create a new event.
	server.POST("/event", middleware.Authenticate, createEvent)

	// Update an event by ID.
	server.PUT("/event/:id", middleware.Authenticate, updateEvent)

	// Delete an event by ID.
	server.DELETE("/event/:id", middleware.Authenticate, deleteEvent)

	// Create a new user.
	server.POST("/user", createUser)

	// Log a user in.
	server.POST("/user/login", loginUser)
}
