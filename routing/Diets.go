package routing

import (
	"github.com/gin-gonic/gin"
	"zakroma_backend/handlers"
	"zakroma_backend/middleware"
)

func DietsRouting(router *gin.RouterGroup) {
	router.GET("/:hash", handlers.GetDietByHash)
	router.POST("/create", middleware.Auth, handlers.CreateDiet)
	router.GET("/current", middleware.Auth, handlers.GetCurrentDiet)
	router.PATCH("/change", middleware.Auth, handlers.ChangeCurrentDiet)
	router.PATCH("/name", middleware.Auth, handlers.ChangeDietName)
	router.GET("/list", middleware.Auth, handlers.GetGroupDiets)
	router.GET("/products", middleware.Auth, handlers.GetDietProducts)
	router.GET("/recipie", middleware.Auth, handlers.GetCurrentDietRecipies)
}
