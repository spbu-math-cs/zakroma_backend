package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"zakroma_backend/middleware"
	"zakroma_backend/routing"
)

func setupSession(router *gin.Engine) {
	cookieStore := cookie.NewStore(secret)
	cookieStore.Options(sessions.Options{MaxAge: 60 * 60 * 24 * 30}) // expire in a month
	router.Use(sessions.Sessions("zakroma_session", cookie.NewStore(secret)))
}

func runHttp(router *gin.Engine) {
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

var secret = []byte("zakrooooooma_baccckendddd_secreeeeet")

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	setupSession(router)

	api := router.Group("/api")
	auth := router.Group("/auth")

	routing.AuthRouting(auth)

	routing.DishesRouting(api.Group("/dishes"))
	routing.ProductsRouting(api.Group("/products"))

	routing.DietsRouting(api.Group("/diets", middleware.Auth))
	routing.DayDietsRouting(api.Group("/diets/day", middleware.Auth))
	routing.MealsRouting(api.Group("/meals", middleware.Auth))

	runHttp(router)
}
