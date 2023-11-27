package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zakroma_backend/stores"
)

func GetMealWithId(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	session := sessions.Default(c)
	userId := session.Get("id")

	if userId == nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	access := stores.CheckMealAccessWithId(id, userId.(int))

	if !access {
		c.Status(http.StatusUnauthorized)
		return
	}

	meal, err := stores.GetMealWithId(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, meal)
}

func CreateMeal(c *gin.Context) {
	type RequestBody struct {
		DietId       int    `json:"diet-id"`
		DayDietIndex int    `json:"day-diet-index"`
		Name         string `json:"name"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	id, err := stores.CreateMeal(requestBody.DietId, requestBody.DayDietIndex, requestBody.Name)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func AddMealDish(c *gin.Context) {
	type RequestBody struct {
		MealId   int `json:"meal-id"`
		DishId   int `json:"dish-id"`
		Portions int `json:"portions"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	err := stores.AddMealDish(requestBody.MealId, requestBody.DishId, requestBody.Portions)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.Status(http.StatusOK)
}
