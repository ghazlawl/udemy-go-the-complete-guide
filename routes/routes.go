package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Get all events.
	server.GET("/events", getEvents)

	// Get an event by ID.
	server.GET("/event/:id", getEvent)

	// Create a new event.
	server.POST("/event", createEvent)

	// Update an event by ID.
	server.PUT("/event/:id", updateEvent)

	// Delete an event by ID.
	server.DELETE("/event/:id", deleteEvent)

	// Create a new user.
	server.POST("/user", createUser)

	// Log a user in.
	server.POST("/user/login", loginUser)
}
