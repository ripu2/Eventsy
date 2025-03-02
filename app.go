package main

import (
	"net/http"

	"example.com/event-management/db"
	"example.com/event-management/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/createEvent", createEvent)
	server.Run(":8080") //localhost:8080

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get events",
			"err":     err.Error(),
		})
		return // return early if there is an error fetching events. No need to continue with the rest of the function.
	}
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
		err = event.Save()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create event",
				"err":     err.Error(),
			})
			return // return early if there is an error saving the event. No need to continue with the rest of the function.
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Event created successfully",
			"event":   event,
		})
	}
}
