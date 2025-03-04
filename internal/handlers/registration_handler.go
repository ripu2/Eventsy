package handlers

import (
	"net/http"
	"strconv"

	"example.com/event-management/internal/models"
	"example.com/event-management/internal/services"
	"example.com/event-management/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetRegistrationHandler(ctx *gin.Context) {
	reg, err := services.GetRegistrations()
	utils.ErrorHandler(ctx, err, "Failed to get events", http.StatusInternalServerError)
	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Fetched Regs Successfully", reg), http.StatusOK)

}

func RegistrationHandler(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	var event models.Event
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	utils.ErrorHandler(ctx, err, "Invalid event ID", http.StatusBadRequest)
	event, err = services.GetEventById(int64(eventId))
	utils.ErrorHandler(ctx, err, "Failed to get event", http.StatusNotFound)
	err = services.RegisterToEvent(&event, userId)
	utils.ErrorHandler(ctx, err, "Failed to Register event", http.StatusInternalServerError)
	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Event created successfully", event), http.StatusCreated)
}

func WithdrawRegistrationHandler(ctx *gin.Context) {

}
