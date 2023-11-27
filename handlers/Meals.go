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
		Dishes       []int  `json:"dishes-ids"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	id, err := stores.CreateMeal(requestBody.DietId, requestBody.DayDietIndex, requestBody.Name, requestBody.Dishes)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
