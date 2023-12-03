package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
)

func MealsRouting(router *gin.RouterGroup) {
	router.GET("/:hash", handlers.GetMealByHash)
	router.POST("/create", handlers.CreateMeal)
	router.POST("/add", handlers.AddMealDish)
}
