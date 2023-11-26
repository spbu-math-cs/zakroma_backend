package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"zakroma_backend/stores"
)

func GetDishWithId(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dish, err := stores.GetDishWithId(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, dish)
}

func GetDishShortWithId(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dish, err := stores.GetDishShortWithId(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, dish)
}

func GetDishesShortWithName(c *gin.Context) {
	type RequestBody struct {
		Name       string `json:"name"`
		RangeBegin int    `json:"range-begin"`
		RangeEnd   int    `json:"range-end"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dishes := stores.GetDishesShortWithName(requestBody.Name, requestBody.RangeBegin, requestBody.RangeEnd)

	c.JSON(http.StatusOK, dishes)
}

func GetDishesShortWithTags(c *gin.Context) {
	tagsStr := c.Params.ByName("tags")
	tags := strings.Split(tagsStr, "$")

	dishes := stores.GetDishesShortWithTags(tags)

	c.JSON(http.StatusOK, dishes)
}
