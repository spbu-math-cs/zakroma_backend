package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
)

func DietsRouting(router *gin.RouterGroup) {
	router.GET("/:id", handlers.GetDietWithId)
	router.GET("/hash/:hash", handlers.GetDietWithHash)
	router.POST("/create", handlers.CreateDiet)
}
