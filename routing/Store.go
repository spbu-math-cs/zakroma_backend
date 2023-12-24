package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"
)

func StoreRouting(router *gin.RouterGroup) {
	router.GET("/", middleware.Auth, handlers.GetGroupStoreList)
	router.POST("/add", middleware.Auth, handlers.AddGroupStoreProduct)
	router.POST("/remove", middleware.Auth, handlers.RemoveGroupStoreProduct)
	router.PATCH("/change", middleware.Auth, handlers.ChangeGroupStoreProduct)
}
