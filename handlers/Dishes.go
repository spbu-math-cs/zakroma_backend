package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zakroma_backend/stores"
)

func GetDishByHash(c *gin.Context) {
	hash := c.Params.ByName("hash")
	if len(hash) == 0 {
		c.String(http.StatusBadRequest, "something bad with field 'hash'")
		return
	}

	dish, err := stores.GetDishByHash(hash)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, dish)
}

func GetDishShortByHash(c *gin.Context) {
	hash := c.Params.ByName("hash")
	if len(hash) == 0 {
		c.String(http.StatusBadRequest, "something bad with field 'hash'")
		return
	}

	dishShort, err := stores.GetDishShortByHash(hash)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, dishShort)
}

func GetDishesShortByName(c *gin.Context) {
	type RequestBody struct {
		Name       string `json:"name"`
		RangeBegin int    `json:"range-begin"`
		RangeEnd   int    `json:"range-end"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	dishes := stores.GetDishesShortByName(requestBody.Name,
		requestBody.RangeBegin, requestBody.RangeEnd)

	c.JSON(http.StatusOK, dishes)
}

func GetDishesShortByTags(c *gin.Context) {
	type RequestBody struct {
		Tags       []string `json:"tags"`
		RangeBegin int      `json:"range-begin"`
		RangeEnd   int      `json:"range-end"`
	}

	var requestBody RequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "request body does not match the protocol")
		return
	}

	dishes := stores.GetDishesShortByTags(requestBody.Tags,
		requestBody.RangeBegin, requestBody.RangeEnd)

	c.JSON(http.StatusOK, dishes)
}
