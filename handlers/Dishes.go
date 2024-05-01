package handlers

import (
	"net/http"
	"zakroma_backend/stores"

	"github.com/gin-gonic/gin"
)

// GetDishByHash godoc
//
// @Tags dishes
// @Accept json
// @Produce json
// @Param hash path string true "Hash блюда"
// @Success 200 {object} schemas.Dish
// @Router /api/dishes/{hash} [get]
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

// GetDishShortByHash godoc
//
// @Description **`products` и `recipe` пустые**
// @Tags dishes
// @Accept json
// @Produce json
// @Param hash path string true "Hash блюда"
// @Success 200 {object} schemas.Dish
// @Router /api/dishes/short/{hash} [get]
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

// GetDishesShortByName godoc
//
// @Description Список базовых блюд, содеражщих в своем названии подстроку *name*
// @Description (выборка с *range-begin* до *range-end*)
// @Tags dishes
// @Accept json
// @Produce json
// @Param data body handlers.GetDishesShortByName.RequestBody true "Тело запроса"
// @Success 200 {array} schemas.Dish
// @Router /api/dishes/name [get]
func GetDishesShortByName(c *gin.Context) {
	type RequestBody struct {
		Name       string `json:"name" example:"салат"`    // Подстрока, которая должна быть в названии блюда
		RangeBegin int    `json:"range-begin" example:"1"` // Начало диапазона (**нумерация с 1**)
		RangeEnd   int    `json:"range-end" example:"5"`   // Конец диапазона
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

// GetDishesShortByTags godoc
//
// @Description Список базовых блюд, обладающих **всеми** тегами из *tags*
// @Description (выборка, с *range-begin* до *range-end*);
// @Description Представлен в виде короткой информации о каждом.
// @Description **`products` и `recipe` пустые**.
// @Tags dishes
// @Accept json
// @Produce json
// @Param data body handlers.GetDishesShortByTags.RequestBody true "Тело запроса"
// @Success 200 {array} schemas.Dish
// @Router /api/dishes/tags [get]
func GetDishesShortByTags(c *gin.Context) {
	type RequestBody struct {
		Tags       []string `json:"tags" example:"breakfast"`
		RangeBegin int      `json:"range-begin" example:"1"` // Начало диапазона (**нумерация с 1**)
		RangeEnd   int      `json:"range-end" example:"5"`   // Конец диапазона
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
