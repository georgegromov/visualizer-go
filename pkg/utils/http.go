package utils

import (
	"errors"
	"net/http"
	"visualizer-go/internal/domains/users"

	"github.com/gin-gonic/gin"
)

// Get user from context
func GetUserFromCtx(ctx *gin.Context) (*users.User, error) {
	uctx, ok := ctx.Get("user")
	if !ok {
		return nil, errors.New("unauthorized: user not found in context")
	}

	user, ok := uctx.(*users.User)
	if !ok {
		return nil, errors.New("unauthorized: invalid user type")
	}

	return user, nil
}

func ReadRequestBody(ctx *gin.Context, request interface{}) error {
	if err := ctx.BindJSON(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request.Context(), request)
}

func CreateTokenCookie() *http.Cookie {
	return &http.Cookie{}
}
