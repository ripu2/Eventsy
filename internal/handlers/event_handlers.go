package handlers

import (
	"net/http"
	"strconv"

	"example.com/event-management/internal/models"
	"example.com/event-management/internal/services"
	"example.com/event-management/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetEventsHandler(ctx *gin.Context) {
	events, err := services.GetEvents()
	ErrorHandler(ctx, err, "Failed to get events", http.StatusInternalServerError)
	HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Fetched Events Successfully", events), http.StatusOK)
}

func CreateEventHandler(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindBodyWithJSON(&event)
	ErrorHandler(ctx, err, "Bad request", http.StatusBadRequest)

	err = services.CreateEvent(event)
	ErrorHandler(ctx, err, "Failed to create event", http.StatusInternalServerError)

	HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Event created successfully", event), http.StatusCreated)
}

func GetEventByIdHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	ErrorHandler(ctx, err, "Invalid event ID", http.StatusBadRequest)

	event, err := services.GetEventById(int64(id))
	ErrorHandler(ctx, err, "Failed to get event", http.StatusNotFound)

	HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Fetched Events Successfully", event), http.StatusOK)

}

func UpdateEventHandler(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindBodyWithJSON(&event)
	ErrorHandler(ctx, err, "Bad request", http.StatusBadRequest)

	idParam, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	ErrorHandler(ctx, err, "Invalid event ID", http.StatusBadRequest)

	err = services.UpdateEvent(event, idParam)
	ErrorHandler(ctx, err, "Failed to update event", http.StatusInternalServerError)

	HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Event updated successfully", event), http.StatusOK)
}

func DeleteEventHandler(ctx *gin.Context) {
	idParam, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	var event models.Event
	ErrorHandler(ctx, err, "Invalid event ID", http.StatusBadRequest)

	err = services.DeleteEvent(event, idParam)
	ErrorHandler(ctx, err, "Failed to get event", http.StatusNotFound)

	HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Event deleted successfully", "Event Deleted !!"), http.StatusOK)

}
