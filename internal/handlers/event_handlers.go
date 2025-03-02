package handlers

import (
	"net/http"

	"example.com/event-management/internal/models"
	"example.com/event-management/internal/services"
	"github.com/gin-gonic/gin"
)

func GetEventsHandler(ctx *gin.Context) {
	events, err := services.GetEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get events",
			"err":     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func CreateEventHandler(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindBodyWithJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"err":     err.Error(),
		})
		return
	}
	err = services.CreateEvent(event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create event",
			"err":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}
