package handlers

import (
	"example.com/event-management/internal/utils"
	"github.com/gin-gonic/gin"
)

type ResponseType struct {
}

func HandleResponse(ctx *gin.Context, responseMap map[string]interface{}, statusCode int) {
	responseBody := gin.H{}
	for key, value := range responseMap {
		responseBody[key] = value
	}
	ctx.JSON(statusCode, responseBody)
}

func ErrorHandler(ctx *gin.Context, err error, errorMessage string, statusCode int) {
	if err != nil {
		HandleResponse(ctx, utils.GenerateMapForResponseType("Error", errorMessage, err.Error()), statusCode)
		panic(err)
	}
}
