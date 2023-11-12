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

func GetDishesShortWithTags(c *gin.Context) {
	tagsStr := c.Params.ByName("tags")
	tags := strings.Split(tagsStr, "$")

	dishes := stores.GetDishesShortWithTags(tags)

	c.JSON(http.StatusOK, dishes)
}
