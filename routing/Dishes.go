package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
)

func DishesRouting(router *gin.RouterGroup) {
	router.GET("/:hash", handlers.GetDishByHash)
	router.GET("/short/:hash", handlers.GetDishShortByHash)
	router.GET("/name", handlers.GetDishesShortByName)
	router.GET("/tags", handlers.GetDishesShortByTags)
}
