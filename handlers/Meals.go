package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zakroma_backend/stores"
)

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

func CreateMeal(c *gin.Context) {
	type RequestBody struct {
		DietHash     string `json:"diet-hash"`
		DayDietIndex int    `json:"day-diet-index"`
		Name         string `json:"name"`
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

	c.JSON(http.StatusOK, gin.H{"hash": hash})
}

func GetAllMealsTags(c *gin.Context) {
	tags, err := stores.GetAllMealsTags()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tags)
}

func AddMealDish(c *gin.Context) {
	type RequestBody struct {
		MealHash string `json:"meal-hash"`
		DishHash string `json:"dish-hash"`
		Portions int    `json:"portions"`
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
