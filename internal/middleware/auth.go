package middleware

import (
	"errors"
	"net/http"

	"example.com/event-management/internal/utils"
	"github.com/gin-gonic/gin"
)

func CheckForAuthentication(ctx *gin.Context) (int64, error) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		utils.ErrorHandler(ctx, errors.New("Unauthenticated"), "Unauthenticated", http.StatusUnauthorized)
		return 0, nil
	}

	id, err := utils.VerifyToken(token)
	if err != nil {
		utils.ErrorHandler(ctx, errors.New(err.Error()), "Unauthenticated", http.StatusUnauthorized)
		return 0, nil
	}

	return id, nil
}
