package templates

import "github.com/gin-gonic/gin"

type Handler interface {
	HandleGet(c *gin.Context)
	HandleGetById(c *gin.Context)
	HandleCreate(c *gin.Context)
	HandleSaveAs(c *gin.Context)
	HandleUpdate(c *gin.Context)
}
