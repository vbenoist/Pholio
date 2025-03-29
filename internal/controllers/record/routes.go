package record

import (
	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/middlewares"
)

func RegisterRoutes(router *gin.Engine) {
	recordSafe := router.Group("/content")
	recordSafe.Use(middlewares.CheckAdmin())
	{
		recordSafe.POST("/record", AddRecord)
		recordSafe.POST("/records", AddRecords)

		recordSafe.PUT("/record/:id", EditRecord)
		recordSafe.DELETE("/record/:id", RemoveRecord)
	}

	// router.PUT("/content/records", EditRecords) // multi or single record edit ?
	// router.DELETE("/content/records", RemoveRecords)
}
