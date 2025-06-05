package diagrams

import "github.com/gin-gonic/gin"

type Handler interface {
	HandleCreate(c *gin.Context)
	HandleGet(c *gin.Context)
	HandleGetById(c *gin.Context)
	HandleUpdate(c *gin.Context)
	HandleDelete(c *gin.Context)
}
