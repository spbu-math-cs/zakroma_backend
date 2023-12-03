package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
)

func ProductsRouting(router *gin.RouterGroup) {
	router.GET("/:id", handlers.GetProductById)
}
