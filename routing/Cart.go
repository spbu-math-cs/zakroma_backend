package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"
)

func CartRouting(router *gin.RouterGroup) {
	router.GET("/", middleware.Auth, handlers.GetGroupCartList)
	router.POST("/add", middleware.Auth, handlers.AddGroupCartProduct)
	router.POST("/remove", middleware.Auth, handlers.RemoveGroupCartProduct)
	router.PATCH("/change", middleware.Auth, handlers.ChangeGroupCartProduct)
}
