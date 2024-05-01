package main

import (
	"zakroma_backend/routing"

	docs "zakroma_backend/docs"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	docs.SwaggerInfo.BasePath = "/"

	router := gin.Default()
	router.Use(gin.Recovery())
	setupSession(router)

	api := router.Group("/api")
	auth := router.Group("/auth")

	routing.AuthRouting(auth)

	routing.DishesRouting(api.Group("/dishes"))
	routing.ProductsRouting(api.Group("/products"))
	routing.DietsRouting(api.Group("/diets"))
	routing.MealsRouting(api.Group("/meals"))
	routing.GroupsRouting(api.Group("/groups"))
	routing.UserRouting(api.Group("/user"))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	runHttp(router)
}
