package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/controllers/image"
	"github.com/jackidu14/pholio/internal/controllers/recently"
	"github.com/jackidu14/pholio/internal/controllers/record"
)

func RegisterRoutes(router *gin.Engine) {
	record.RegisterRoutes(router)
	recently.RegisterRoutes(router)
	image.RegisterRoutes(router)
}
