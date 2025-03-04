package handlers

import (
	"net/http"

	"example.com/event-management/internal/models"
	"example.com/event-management/internal/services"
	"example.com/event-management/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetUsersEventHandler(ctx *gin.Context) {
	users, err := services.GetUsers()
	utils.ErrorHandler(ctx, err, "Failed to get users", http.StatusInternalServerError)
	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data", "Fetched Users Successfully", users), http.StatusOK)
}

func CreateUserHandler(ctx *gin.Context) {
	var user *models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ErrorHandler(ctx, err, "Failed to bind JSON", http.StatusBadRequest)
		return
	}

	if err := services.CreateUser(user); err != nil {
		utils.ErrorHandler(ctx, err, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("data", "User created successfully", user), http.StatusCreated)
}

func LoginHandler(ctx *gin.Context) {
	var LoginRequest models.LoginRequest
	err := ctx.ShouldBindJSON(&LoginRequest)
	utils.ErrorHandler(ctx, err, "Failed to bind JSON", http.StatusBadRequest)

	token, err := services.LoginUser(LoginRequest)
	utils.ErrorHandler(ctx, err, "Invalid credentials", http.StatusUnauthorized)

	utils.HandleResponse(ctx, utils.GenerateMapForResponseType("token", "Authenticated", token), http.StatusAccepted)

}
