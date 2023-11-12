package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
)

func DishesRouting(router *gin.RouterGroup) {
	router.GET("/:id", handlers.GetDishWithId)
	router.GET("/short/:id", handlers.GetDishShortWithId)
	router.GET("/tags/:tags", handlers.GetDishesShortWithTags)
}
