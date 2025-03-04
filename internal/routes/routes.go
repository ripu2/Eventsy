package routes

import (
	"example.com/event-management/internal/handlers"
	"example.com/event-management/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {
	server.GET("/events", handlers.GetEventsHandler)

	authServer := server.Group("/")
	authServer.Use(middleware.CheckForAuthentication)
	authServer.POST("/createEvent", handlers.CreateEventHandler)
	authServer.GET("/events/:id", handlers.GetEventByIdHandler)
	authServer.PUT("/updateEvent/:id", handlers.UpdateEventHandler)
	authServer.DELETE("/deleteEvent/:id", handlers.DeleteEventHandler)

	// User routes

	server.POST("/createUser", handlers.CreateUserHandler)
	server.GET("/users", handlers.GetUsersEventHandler)
	server.POST("/login", handlers.LoginHandler)
}
