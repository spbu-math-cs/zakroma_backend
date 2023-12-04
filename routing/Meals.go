package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"
)

func MealsRouting(router *gin.RouterGroup) {
	router.GET("/:hash", handlers.GetMealByHash)
	router.POST("/create", middleware.Auth, handlers.CreateMeal)
	router.POST("/add", middleware.Auth, handlers.AddMealDish)
}
