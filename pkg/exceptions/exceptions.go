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
	ErrBadRequestString          = "Bad request"
	ErrInvalidCredentialsString  = "Invalid credentials"
	ErrInvalidQueryParamsString  = "Invalid query params"
	ErrNotFoundString            = "Not found"
	ErrUnauthorizedString        = "Unauthorized"
	ErrForbiddenString           = "Forbidden"
	ErrUsernameTakenString       = "Username already taken"
	ErrInternalServerErrorString = "Internal server error"
	ErrNoCookieString            = "Cookie not found"
	ErrRequestTimeoutString      = "Request timeout"
	ErrInvalidAuthTokenString    = "Invalid authorization token"
)

var (
	ErrBadRequest          = errors.New(ErrBadRequestString)
	ErrInvalidCredentials  = errors.New(ErrInvalidCredentialsString)
	ErrInvalidQueryParams  = errors.New(ErrInvalidQueryParamsString)
	ErrNotFound            = errors.New(ErrNotFoundString)
	ErrUnauthorized        = errors.New(ErrUnauthorizedString)
	ErrForbidden           = errors.New(ErrForbiddenString)
	ErrUsernameTaken       = errors.New(ErrUsernameTakenString)
	ErrInternalServerError = errors.New(ErrInternalServerErrorString)
	ErrNoCookie            = errors.New(ErrNoCookieString)
	ErrRequestTimeoutError = errors.New(ErrRequestTimeoutString)
	ErrInvalidJWTToken     = errors.New(ErrInvalidAuthTokenString)
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
		ExcError:  ErrBadRequest.Error(),
		ExcCauses: causes,
	}
}

func NewNotFoundError(causes interface{}) Exception {
	return &HttpException{
		ExcStatus: http.StatusNotFound,
		ExcError:  ErrNotFound.Error(),
		ExcCauses: causes,
	}
}

func NewUnauthorizedError(causes interface{}) Exception {
	return &HttpException{
		ExcStatus: http.StatusUnauthorized,
		ExcError:  ErrUnauthorized.Error(),
		ExcCauses: causes,
	}
}

func NewForbiddenError(causes interface{}) Exception {
	return &HttpException{
		ExcStatus: http.StatusForbidden,
		ExcError:  ErrForbidden.Error(),
		ExcCauses: causes,
	}
}

func NewInternalServerError(causes interface{}) Exception {
	result := &HttpException{
		ExcStatus: http.StatusInternalServerError,
		ExcError:  ErrInternalServerError.Error(),
		ExcCauses: causes,
	}
	return result
}

func parseError(err error) Exception {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewHttpException(http.StatusNotFound, ErrNotFound.Error(), err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewHttpException(http.StatusRequestTimeout, ErrRequestTimeoutError.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "cookie"):
		return NewHttpException(http.StatusUnauthorized, ErrUnauthorized.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "token"):
		return NewHttpException(http.StatusUnauthorized, ErrUnauthorized.Error(), err)

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
