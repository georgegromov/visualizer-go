package exceptions

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	ErrBadRequest          = "Bad request"
	ErrInvalidCredentials  = "Invalid credentials"
	ErrNotFound            = "Not found"
	ErrUnauthorized        = "Unauthorized"
	ErrForbidden           = "Forbidden"
	ErrUsernameTaken       = "Username already taken"
	ErrInternalServerError = "Internal server error"
	ErrNoCookie            = "Cookie not found"
	ErrRequestTimeout      = "Request timeout"
	ErrInvalidAuthToken    = "Invalid authorization token"
)

var (
	BadRequest          = errors.New(ErrBadRequest)
	InvalidCredentials  = errors.New(ErrInvalidCredentials)
	NotFound            = errors.New(ErrNotFound)
	Unauthorized        = errors.New(ErrUnauthorized)
	Forbidden           = errors.New(ErrForbidden)
	UsernameTaken       = errors.New(ErrUsernameTaken)
	InternalServerError = errors.New(ErrInternalServerError)
	NoCookie            = errors.New(ErrNoCookie)
	RequestTimeoutError = errors.New(ErrRequestTimeout)
	InvalidJWTToken     = errors.New(ErrInvalidAuthToken)
)

type Exception interface {
	Status() int
	Error() string
	Causes() interface{}
}

type HttpException struct {
	ExcStatus int         `json:"status,omitempty"`
	ExcError  string      `json:"error,omitempty"`
	ExcCauses interface{} `json:"-"`
}

func (e *HttpException) Status() int {
	return e.ExcStatus
}

func (e *HttpException) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ExcStatus, e.ExcError, e.ExcCauses)
}

func (e *HttpException) Causes() interface{} {
	return e.ExcCauses
}

func NewHttpException(status int, err string, causes interface{}) Exception {
	return &HttpException{
		ExcStatus: status,
		ExcError:  err,
		ExcCauses: causes,
	}
}

func NewBadRequestError(causes interface{}) Exception {
	return &HttpException{
		ExcStatus: http.StatusBadRequest,
		ExcError:  BadRequest.Error(),
		ExcCauses: causes,
	}
}

func NewNotFoundError(causes interface{}) Exception {
	return &HttpException{
		ExcStatus: http.StatusNotFound,
		ExcError:  NotFound.Error(),
		ExcCauses: causes,
	}
}

func NewUnauthorizedError(causes interface{}) Exception {
	return &HttpException{
		ExcStatus: http.StatusUnauthorized,
		ExcError:  Unauthorized.Error(),
		ExcCauses: causes,
	}
}

func NewForbiddenError(causes interface{}) Exception {
	return &HttpException{
		ExcStatus: http.StatusForbidden,
		ExcError:  Forbidden.Error(),
		ExcCauses: causes,
	}
}

func NewInternalServerError(causes interface{}) Exception {
	result := &HttpException{
		ExcStatus: http.StatusInternalServerError,
		ExcError:  InternalServerError.Error(),
		ExcCauses: causes,
	}
	return result
}

func parseError(err error) Exception {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewHttpException(http.StatusNotFound, NotFound.Error(), err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewHttpException(http.StatusRequestTimeout, RequestTimeoutError.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "cookie"):
		return NewHttpException(http.StatusUnauthorized, Unauthorized.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "token"):
		return NewHttpException(http.StatusUnauthorized, Unauthorized.Error(), err)

	default:
		if exception, ok := err.(Exception); ok {
			return exception
		}
		return NewInternalServerError(err)
	}
}

func SendErrorResponse(err error) (int, interface{}) {
	return parseError(err).Status(), parseError(err)
}
