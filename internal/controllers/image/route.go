package image

import (
	"github.com/gin-gonic/gin"
	"github.com/vbenoist/pholio/internal/middlewares"
)

func RegisterRoutes(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	recordSafe := router.Group("/content/record/:id/image")
	recordSafe.Use(middlewares.CheckRecordId())
	{
		recordSafe.GET("/thumb", GetThumbImage)
		recordSafe.GET("/mid", GetMidImage)
		recordSafe.GET("/orig", GetOrigImage)

		recordSafe.POST("", AddImage)
	}
}
