package canvases

import "github.com/gin-gonic/gin"

type Handler interface {
	HandleCreate(c *gin.Context)
	HandleGetByTemplateId(c *gin.Context)
	HandleUpdate(c *gin.Context)
	HandleDelete(c *gin.Context)
}
