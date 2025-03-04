package middleware

import (
	"errors"
	"net/http"

	"example.com/event-management/internal/utils"
	"github.com/gin-gonic/gin"
)

func CheckForAuthentication(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		utils.ErrorHandler(ctx, errors.New("Unauthenticated"), "Unauthenticated", http.StatusUnauthorized)
	}

	id, err := utils.VerifyToken(token)
	if err != nil {
		utils.ErrorHandler(ctx, errors.New(err.Error()), "Unauthenticated", http.StatusUnauthorized)
	}
	ctx.Set("userId", id)
	ctx.Next()

}
