package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"
)

func GroupsRouting(router *gin.RouterGroup) {
	router.POST("/create", middleware.Auth, handlers.CreateGroup)
	router.POST("/user/add", middleware.Auth, handlers.AddGroupUser)
	router.PATCH("/role", middleware.Auth, handlers.ChangeRole)
	router.POST("/diet/add", middleware.Auth, handlers.AddGroupDiet)
	router.GET("/list", middleware.Auth, handlers.GetAllUserGroups)
	router.PATCH("/change", middleware.Auth, handlers.ChangeCurrentGroup)
}
