package record

import (
	"github.com/gin-gonic/gin"
	"github.com/vbenoist/pholio/internal/middlewares"
)

func RegisterRoutes(router *gin.Engine) {
	adminSafe := router.Group("/content")
	adminSafe.Use(middlewares.CheckAdmin())
	{
		adminSafe.POST("/record", AddRecord)
		adminSafe.POST("/records", AddRecords)

		adminSafe.PUT("/record/:id", EditRecord)
		adminSafe.DELETE("/record/:id", RemoveRecord)

		adminSafe.GET("/records/detailed", GetDetailedRecords)
	}

	router.GET("/content/records/recently", GetRecentlyContent)
	router.GET("/content/records/per-date", GetRecordsPerDate)
	router.GET("/content/records/per-location", GetRecordsPerLocation)

	// router.PUT("/content/records", EditRecords) // multi or single record edit ?
	// router.DELETE("/content/records", RemoveRecords)
}
