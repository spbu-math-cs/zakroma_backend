package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func GetDishesShortWithTags(c *gin.Context) {
	type RequestBody struct {
		Tags       []string `json:"tags"`
		RangeBegin int      `json:"range-begin"`
		RangeEnd   int      `json:"range-end"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	tags := requestBody.Tags
	rangeBegin := requestBody.RangeBegin
	rangeEnd := requestBody.RangeEnd

	dishes := stores.GetDishesShortWithTags(tags, rangeBegin, rangeEnd)

	c.JSON(http.StatusOK, dishes)
}
