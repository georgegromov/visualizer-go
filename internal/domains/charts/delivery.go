package charts

import "github.com/gin-gonic/gin"

type Handler interface {
  HandleGetByID(c *gin.Context)
	HandleGetByCanvasId(c *gin.Context)
	HandleCreate(c *gin.Context)
	HandleUpdate(c *gin.Context)
	HandleDelete(c *gin.Context)
}
