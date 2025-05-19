package measurements

import "github.com/gin-gonic/gin"

type Handler interface {
	HandleGetByChartID(c *gin.Context)
	HandleCreate(c *gin.Context)
	HandleUpdate(c *gin.Context)
	HandleDelete(c *gin.Context)
}
