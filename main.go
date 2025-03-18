package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackidu14/pholio/internal/database/connector"
	"github.com/jackidu14/pholio/internal/helpers/cfg"
	"github.com/jackidu14/pholio/internal/server"
)

func setupRouter(config cfg.ServerConfig) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Front.Url}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
	router.Use(cors.New(corsConfig))

	server.RegisterRoutes(router)
	return router
}

func main() {
	config := cfg.SetServerConfig()
	connector.Connect()
	defer connector.Disconnect()

	router := setupRouter(config)
	router.Run(fmt.Sprintf(":%s", config.Instance.Port))
}
