package main

import (
	"net/http"

	"example.com/event-management/db"
	"example.com/event-management/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Run(":8080") //localhost:8080
	isDbReady := make(chan bool)
	go db.InitDB(isDbReady)

	if <-isDbReady {
		server.GET("/events", getEvents)
		server.POST("/createEvent", createEvent)

	}
}

func getEvents(context *gin.Context) {
	events, _ := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindBodyWithJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"err":     err.Error(),
		})
	} else {
		event.ID = 1
		event.UserID = 1
		event.Save()
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Event created successfully",
			"event":   event,
		})
	}
}
