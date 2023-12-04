package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"
)

func DietsRouting(router *gin.RouterGroup) {
	router.GET("/:hash", handlers.GetDietByHash)
	router.POST("/create", middleware.Auth, handlers.CreateDiet)
	router.GET("/current", middleware.Auth, handlers.GetCurrentDiet)
}
