package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"visualizer-go/internal/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(log *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			log.Error("Authorization header is missing")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			ctx.Abort()
			return
		}

		expectedToken := "test_token_123"

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token != expectedToken {
			log.Error(fmt.Sprintf("Invalid token: %s", token))
			response.Error(ctx, http.StatusUnauthorized, "unauthorized", gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
