package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
)

func DietsRouting(router *gin.RouterGroup) {
	router.GET("/:id", handlers.GetDietWithId)
}
