package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vbenoist/pholio/internal/middlewares"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/auth/login", Login)

	authCheck := router.Group("/auth")
	authCheck.Use(middlewares.CheckAdmin())
	{
		/* Just simple 200 status returned if user is sent authenticated request */
		authCheck.GET("/check", Check)
	}
}
