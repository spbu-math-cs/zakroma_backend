package routing

import (
	"zakroma_backend/handlers"

	"github.com/gin-gonic/gin"
)

func UserRouting(router *gin.RouterGroup) {
	router.GET("/name", handlers.GetName)
}
