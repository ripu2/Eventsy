package routes

import (
	"example.com/event-management/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {
	server.GET("/events", handlers.GetEventsHandler)
	server.POST("/createEvent", handlers.CreateEventHandler)
	server.GET("/events/:id", handlers.GetEventByIdHandler)
	server.PUT("/updateEvent/:id", handlers.UpdateEvent)
}
