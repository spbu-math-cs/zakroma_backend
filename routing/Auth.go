package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"
)

func AuthRouting(router *gin.RouterGroup) {
	router.GET("/ping", middleware.Auth, handlers.Ping)
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
}
