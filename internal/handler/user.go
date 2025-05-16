package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"
	"visualizer-go/internal/repository"
	"visualizer-go/internal/response"
	"visualizer-go/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrUserIDMissing          = errors.New("user ID is missing")
	ErrInvalidUserIDFormat    = errors.New("invalid user ID format")
	ErrFailedToCreateUser     = errors.New("failed to create user")
	ErrFailedToUpdateUser     = errors.New("failed to update user")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserInvalidRequestData = errors.New("invalid user request data")
)

// loginUser godoc
//
// @Summary Login user
// @Tags Auth
// @Router /auth/login [post]
func (h *Handler) handleLogin(ctx *gin.Context) {
	const op = "handler.login"

	var userLoginDto dto.UserLoginDto
	if err := ctx.ShouldBind(&userLoginDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, ErrUserInvalidRequestData.Error(), err)
		return
	}

	userWithToken, err := h.services.Login(ctx.Request.Context(), userLoginDto)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, err.Error(), repository.ErrInvalidCredentials.Error())
		return
	}

	// TODO encapsulate in CreateTokenCookie func
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    userWithToken.Token,
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(7 * 24 * time.Hour), // 7 дней
		MaxAge:   3600 * 24 * 7,                      // В секундах (7 дней)
		HttpOnly: true,                               // Защита от доступа из JS
		Secure:   true,                               // Только HTTPS
		SameSite: http.SameSiteNoneMode,              // Запрещает кросс-доменные запросы с куки
	})

	response.Success(ctx, http.StatusOK, "Logged in successfully", gin.H{
		"user":        userWithToken.User,
		"accessToken": userWithToken.Token,
	})
}

// getUserById godoc
//
// @Summary Get user by id
// @Tags Users
// @Router /users/{id} [get]
func (h *Handler) handleGetUserByID(ctx *gin.Context) {
	const op = "handler.getUserByID"

	userIDStr := ctx.Param("id")
	if userIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrUserIDMissing))
		response.Error(ctx, http.StatusBadRequest, ErrUserIDMissing.Error(), ErrUserIDMissing.Error())
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, ErrInvalidUserIDFormat.Error(), ErrInvalidUserIDFormat.Error())
		return
	}

	user, err := h.services.User.GetByID(ctx, userID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusInternalServerError, ErrUserNotFound.Error(), ErrUserNotFound.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "User fetched successfully", user)
}

// createUser godoc
//
// @Summary Create user
// @Tags Users
// @Router /users [post]
func (h *Handler) handleCreateUser(ctx *gin.Context) {
	const op = "handler.Handler.createUser"

	user := &models.User{}
	if err := utils.ReadRequestBody(ctx, user); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, ErrUserInvalidRequestData.Error(), err)
		return
	}

	if err := h.services.User.Create(ctx.Request.Context(), user); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusInternalServerError, ErrFailedToCreateUser.Error(), err)
		return
	}

	response.Success(ctx, http.StatusCreated, "User created successfully", nil)
}

// updateUser godoc
//
// @Summary Update user
// @Tags Users
// @Router /users/{id} [patch]
func (h *Handler) handleUpdateUser(ctx *gin.Context) {
	const op = "handler.Handler.updateUser"

	userIDStr := ctx.Param("id")
	if userIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrUserIDMissing))
		response.Error(ctx, http.StatusBadRequest, ErrUserIDMissing.Error(), nil)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, ErrInvalidUserIDFormat.Error(), err)
		return
	}

	var userUpdateDto dto.UserUpdateDto
	if err = ctx.ShouldBindJSON(&userUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, ErrUserInvalidRequestData.Error(), err)
		return
	}

	if err = h.services.User.Update(ctx.Request.Context(), userID, userUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusInternalServerError, ErrFailedToUpdateUser.Error(), err)
		return
	}

	response.Success(ctx, http.StatusOK, "User updated successfully", nil)
}
