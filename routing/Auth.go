package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
)

func AuthRouting(router *gin.RouterGroup) {
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
}
