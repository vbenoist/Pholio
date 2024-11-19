package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/pkg/recently"
)

func RegisterRoutes(router *gin.Engine) {
	recently.RegisterRoutes(router)
}
