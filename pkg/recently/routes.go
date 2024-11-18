package recently

import (
	"github.com/gin-gonic/gin"
)

func registerRecentlyRoutes(router *gin.Engine) {
	router.GET("/content/recently", getRecentlyContent)
}
