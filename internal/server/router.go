package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/controllers/auth"
	"github.com/jackidu14/pholio/internal/controllers/image"
	"github.com/jackidu14/pholio/internal/controllers/record"
	"github.com/jackidu14/pholio/internal/helpers/cfg"
)

func SetupRouter(config cfg.ServerConfig) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Front.Url}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	RegisterRoutes(router)
	return router
}

func RegisterRoutes(router *gin.Engine) {
	auth.RegisterRoutes(router)
	record.RegisterRoutes(router)
	image.RegisterRoutes(router)
}
