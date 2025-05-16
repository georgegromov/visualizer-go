package middlewares

import (
	"log/slog"
	"net/http"
	"strings"
	"visualizer-go/internal/response"
	"visualizer-go/internal/service"
	jwt_manager "visualizer-go/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware(log *slog.Logger, services *service.Service, jwtManager *jwt_manager.JwtManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		const op = "middleware.AuthMiddleware"

		var accessToken string
		cookieAccessToken, err := ctx.Cookie("accessToken")

		if cookieAccessToken == "" || err != nil {
			log.Info("no access token set in cookie")

			authHeader := ctx.GetHeader("Authorization")
			if authHeader == "" {
				response.Error(ctx, http.StatusUnauthorized, "User unauthorized", err.Error())
				ctx.Abort()
				return
			}

			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				slog.Warn("invalid authorization header")
				response.Error(ctx, http.StatusUnauthorized, "Invalid authorization header", err.Error())
				ctx.Abort()
				return
			}

			accessToken = headerParts[1]
		} else {
			log.Info("using access token set in cookie")
			accessToken = cookieAccessToken
		}

		claims, err := jwtManager.ParseToken(accessToken)
		if err != nil {
			log.Info("an error occured while parsing token:", slog.String("error", err.Error()))
			// TODO encapsulate in DeleteTokenCookie func
			ctx.SetCookie("accessToken", "", -1, "/", "", true, true)
			response.Error(ctx, http.StatusUnauthorized, "Invalid authorization header", err.Error())
			ctx.Abort()
			return
		}

		cuid := claims.UserID.String()

		uid, err := uuid.Parse(cuid)
		if err != nil {
			log.Info("an error occured while parsing uid from token claims:", slog.String("error", err.Error()))
			ctx.SetCookie("accessToken", "", -1, "/", "", true, true)
			response.Error(ctx, http.StatusUnauthorized, "Invalid user id", err.Error())
			ctx.Abort()
			return
		}

		user, err := services.User.GetByID(ctx, uid)
		if err != nil {
			log.Info("an error occured while getting user:", slog.String("error", err.Error()))
			ctx.SetCookie("accessToken", "", -1, "/", "", true, true)
			response.Error(ctx, http.StatusUnauthorized, "User not found", err.Error())
			ctx.Abort()
			return
		}

		ctx.Set("user", &user)

		ctx.Next()
	}
}
