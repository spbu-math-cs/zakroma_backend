package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
)

func DayDietsRouting(router *gin.RouterGroup) {
	router.GET("/:id", handlers.GetDayDietWithId)
}
