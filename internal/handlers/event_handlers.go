package handlers

import (
	"net/http"
	"strconv"

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

func GetEventByIdHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"err":     err.Error(),
		})
		return
	}

	event, err := services.GetEventById(int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to get event",
			"err":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event found",
		"event":   event,
		"id":      idParam,
	})
}

func UpdateEvent(ctx *gin.Context) {
	var event models.Event
	error := ctx.ShouldBindBodyWithJSON(&event)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update event",
			"err":     error.Error(),
		})
	}
	idParam, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
			"err":     err.Error(),
		})
	}

	err = services.UpdateEvent(event, idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update event",
			"err":     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   event,
		"id":      idParam,
	})
}
