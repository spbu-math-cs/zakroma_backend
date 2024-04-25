package handlers

import (
	"net/http"
	"zakroma_backend/stores"

	"github.com/gin-gonic/gin"
)

// GetMealByHash godoc
//
// @Tags meals
// @Accept json
// @Produce json
// @Param hash path string true "Hash приема пищи"
// @Success 200 {object} schemas.Meal
// @Router /api/meals/{hash} [get]
func GetMealByHash(c *gin.Context) {
	hash := c.Params.ByName("hash")
	if len(hash) == 0 {
		c.String(http.StatusBadRequest, "something bad with field 'hash'")
		return
	}

	meal, err := stores.GetMealByHash(hash)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, meal)
}

// CreateMeal godoc
//
// @Tags meals
// @Accept json
// @Produce json
// @Param data body handlers.CreateMeal.RequestBody true "Тело запроса"
// @Success 200 {object} handlers.CreateMeal.ResponseBody
// @Security Bearer
// @Router /api/meals/create [post]
func CreateMeal(c *gin.Context) {
	type RequestBody struct {
		DietHash     string `json:"diet-hash" example:"92bc3119092103d17059ba75ca19db9541d282e929c43cbb72de1231429d862d"`
		DayDietIndex int    `json:"day-diet-index" example:"3"`
		Name         string `json:"name" example:"завтрак в четверг"`
	}
	type ResponseBody struct {
		Hash string `json:"hash" example:"5227edc70e407c980ff906f7b0dc241cf3fa2b8139348c382b33818a49e3b36a"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	hash, err := stores.CreateMeal(requestBody.DietHash,
		requestBody.DayDietIndex, requestBody.Name)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ResponseBody{Hash: hash})
}

// GetAllMealsTags godoc
//
// @Tags meals
// @Accept json
// @Produce json
// @Success 200 {array} schemas.Tag
// @Security Bearer
// @Router /api/meals/tags [get]
func GetAllMealsTags(c *gin.Context) {
	tags, err := stores.GetAllMealsTags()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tags)
}

// AddMealDish godoc
//
// @Tags meals
// @Accept json
// @Produce json
// @Param data body handlers.AddMealDish.RequestBody true "Тело запроса"
// @Success 200
// @Security Bearer
// @Router /api/meals/add [post]
func AddMealDish(c *gin.Context) {
	type RequestBody struct {
		MealHash string `json:"meal-hash" example:"f14a528413bc023996568dceaf09295b2b680937d89ab278eaab75551428be52"`
		DishHash string `json:"dish-hash" example:"43cb59219f2c150a5c9cf84eafadf1fc2262cbcd1ec13ddf43ce653da9339784"`
		Portions int    `json:"portions" example:"1"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	err := stores.AddMealDish(requestBody.MealHash,
		requestBody.DishHash, requestBody.Portions)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
