package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/lib/response"
)

var (
	ErrUserIDMissing          = errors.New("user ID is missing")
	ErrInvalidUserIDFormat    = errors.New("invalid user ID format")
	ErrFailedToCreateUser     = errors.New("failed to create user")
	ErrFailedToUpdateUser     = errors.New("failed to update user")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserInvalidRequestData = errors.New("invalid user request data")
)

func (h *Handler) login(ctx *gin.Context) {
	const op = "handler.login"

	var userLoginDto dto.UserLoginDto
	if err := ctx.ShouldBind(&userLoginDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, ErrUserInvalidRequestData.Error(), err)
		return
	}

	user, token, err := h.services.Login(ctx.Request.Context(), userLoginDto)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	response.Success(ctx, http.StatusOK, "Logged in successfully", gin.H{
		"user":  user,
		"token": token,
	})
}

func (h *Handler) createUser(ctx *gin.Context) {
	const op = "handler.Handler.createUser"

	var userCreateDto dto.UserCreateDto
	if err := ctx.ShouldBindJSON(&userCreateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, ErrUserInvalidRequestData.Error(), err)
		return
	}

	if err := h.services.User.Create(ctx.Request.Context(), userCreateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusInternalServerError, ErrFailedToCreateUser.Error(), err)
		return
	}

	response.Success(ctx, http.StatusCreated, "User created successfully", nil)
}

func (h *Handler) updateUser(ctx *gin.Context) {
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
