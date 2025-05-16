package users

import "github.com/gin-gonic/gin"

type Handler interface {
	HandleLogin(c *gin.Context)
	HandleGetById(c *gin.Context)
	HandleCreate(c *gin.Context)
	HandleUpdate(c *gin.Context)
}
