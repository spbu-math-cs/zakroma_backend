package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
)

func MealsRouting(router *gin.RouterGroup) {
	router.GET("/:id", handlers.GetMealWithId)
}
