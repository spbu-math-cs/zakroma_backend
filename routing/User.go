package routing

import (
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouting(router *gin.RouterGroup) {
	router.GET("/name", middleware.Auth, handlers.GetName)
}
