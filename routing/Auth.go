package routing

import (
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouting(router *gin.RouterGroup) {
	router.GET("/ping", middleware.Auth, handlers.Ping)
	router.GET("/name", handlers.GetName)
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
}
