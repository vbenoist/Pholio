package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/pkg/recently"
	"github.com/jackidu14/pholio/pkg/record"
)

func RegisterRoutes(router *gin.Engine) {
	record.RegisterRoutes(router)
	recently.RegisterRoutes(router)
}
