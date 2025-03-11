package record

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/content/record", AddRecord)
	router.POST("/content/records", AddRecords)

	router.PUT("/content/record/:id", EditRecord)
	router.DELETE("/content/record/:id", RemoveRecord)

	// router.PUT("/content/records", EditRecords) // multi or single record edit ?
	// router.DELETE("/content/records", RemoveRecords)
}
