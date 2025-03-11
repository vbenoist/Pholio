package image

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/content/record/:id/image", AddImage)
}
