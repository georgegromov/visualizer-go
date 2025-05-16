package utils

import (
	"github.com/gin-gonic/gin"
)

func ReadRequestBody(ctx *gin.Context, request interface{}) error {
	if err := ctx.BindJSON(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request.Context(), request)
}
