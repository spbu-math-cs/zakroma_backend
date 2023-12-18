package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"
)

func GroupsRouting(router *gin.RouterGroup) {
	router.GET("/list", middleware.Auth, handlers.GetAllUserGroups)
}
