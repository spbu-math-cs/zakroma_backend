package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"
)

func MealsRouting(router *gin.RouterGroup) {
	router.GET("/:hash", handlers.GetMealByHash)
	router.POST("/create/name", middleware.Auth, handlers.CreateMealByName)
	router.GET("/tags", middleware.Auth, handlers.GetAllMealsTags)
	router.POST("/create/tag", middleware.Auth, handlers.CreateMealByTag)
	router.POST("/add", middleware.Auth, handlers.AddMealDish)
}
