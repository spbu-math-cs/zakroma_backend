package routing

import (
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"

	"github.com/gin-gonic/gin"
)

func DietsRouting(router *gin.RouterGroup) {
	router.GET("/:hash", handlers.GetDietByHash)
	router.POST("/create", middleware.Auth, handlers.CreateDiet)
	router.GET("/current", middleware.Auth, handlers.GetCurrentDiet)
	router.PATCH("/change", middleware.Auth, handlers.ChangeCurrentDiet)
	router.PATCH("/name", middleware.Auth, handlers.ChangeDietName)
	router.GET("/list", middleware.Auth, handlers.GetGroupDiets)
	router.POST("/products", middleware.Auth, handlers.GetDietProducts)
	router.GET("/recipie", middleware.Auth, handlers.GetCurrentDietRecipies)
}
